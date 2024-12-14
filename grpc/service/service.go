package service

import (
	"bytes"
	"context"
	"image"
	"io"
	"sandbox/grpc/gen/pkg/v1"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/mostynb/go-grpc-compression/snappy"
	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/status"
)

type Service struct {
	pkg.UnimplementedServiceServer
}

func NewService() *Service {
	return &Service{}
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
		if err == io.EOF {
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
	return nil
}
