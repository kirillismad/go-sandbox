package service

import (
	"bytes"
	"context"
	"errors"
	"image"
	"io"
	"sandbox/grpc/gen/pkg/v1"
	"strconv"
	"sync"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/mostynb/go-grpc-compression/snappy"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type message struct {
	text   string
	sender int64
}

type Service struct {
	pkg.UnimplementedServiceServer
	chatters map[int64]chan message
	mu       sync.RWMutex
}

func NewService() pkg.ServiceServer {
	return &Service{
		chatters: make(map[int64]chan message),
	}
}

// unary
func (s *Service) Get(ctx context.Context, request *pkg.GetRequest) (*pkg.GetResponse, error) {
	return &pkg.GetResponse{
		Id:   request.GetId(),
		Text: gofakeit.Phrase(),
	}, nil
}

// client streaming
func (s *Service) Download(request *pkg.DownloadRequest, stream pkg.Service_DownloadServer) error {
	content := gofakeit.ImageJpeg(int(request.GetWidth()), int(request.GetHeight()))
	const chunkSize = 512

	for i := 0; i < len(content); i += chunkSize {
		chunk := content[i:min(i+chunkSize, len(content))]
		err := stream.Send(&pkg.DownloadResponse{Chunk: chunk})
		if err != nil {
			return status.Errorf(codes.Internal, "failed to send chunk: %v", err)
		}
	}
	return nil
}

// server streaming
func (s *Service) Upload(stream pkg.Service_UploadServer) error {
	var content []byte
	for {
		req, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "failed to receive chunk: %v", err)
		}
		content = append(content, req.GetChunk()...)
	}
	_, format, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "failed to decode image: %v", err)
	}
	if format != "jpeg" {
		return status.Errorf(codes.InvalidArgument, "invalid image format: %s", format)
	}
	return stream.SendAndClose(&pkg.UploadResponse{Size: int64(len(content))})
}

// bidirectional streaming
func (s *Service) Chat(stream pkg.Service_ChatServer) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.InvalidArgument, "missing metadata")
	}
	if len(md.Get("user_id")) == 0 || len(md.Get("user_id")) > 1 {
		return status.Error(codes.InvalidArgument, "user_id must be provided once")
	}
	userId, err := strconv.ParseInt(md.Get("user_id")[0], 10, 64)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid user_id: %v", err)
	}
	s.mu.Lock()
	s.chatters[userId] = make(chan message)
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		close(s.chatters[userId])
		delete(s.chatters, userId)
	}()

	reader := lo.Async(func() error {
		for {
			select {
			case msg, ok := <-s.chatters[userId]:
				if !ok {
					return nil
				}
				err := stream.Send(&pkg.MessageResponse{Text: msg.text, Sender: msg.sender})
				if err != nil {
					return status.Errorf(codes.Internal, "failed to send message: %v", err)
				}
			case <-ctx.Done():
				return nil
			}
		}
	})

	writer := lo.Async(func() error {
		for {
			msg, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return nil
			}
			if err != nil {
				return status.Errorf(codes.Internal, "failed to receive message: %v", err)
			}
			s.mu.RLock()
			receiver, ok := s.chatters[msg.GetReciever()]
			s.mu.RUnlock()
			if !ok {
				return status.Errorf(codes.NotFound, "receiver not found: %d", msg.GetReciever())
			}

			select {
			case receiver <- message{text: msg.GetText(), sender: userId}:
			case <-ctx.Done():
				return nil
			}
		}
	})

	select {
	case err := <-reader:
		return err
	case err := <-writer:
		return err
	}
}
