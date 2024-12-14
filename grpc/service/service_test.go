package service

import (
	"bytes"
	"context"
	"errors"
	"image"
	"io"
	"log"
	"net"
	"sandbox/grpc/gen/pkg/v1"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/mostynb/go-grpc-compression/snappy"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc/encoding/gzip"
)

type GrpcServiceSuite struct {
	suite.Suite
	client pkg.ServiceClient
}

func (s *GrpcServiceSuite) SetupSuite() {
	// setup server
	for _, address := range servers {
		lis, err := net.Listen("tcp", address)
		s.Require().NoError(err)

		srv := grpc.NewServer()
		pkg.RegisterServiceServer(srv, NewService())

		go func() {
			if err := srv.Serve(lis); err != nil {
				s.Fail("failed to serve: %v", err)
				log.Fatal()
			}
		}()
		s.T().Cleanup(func() {
			srv.Stop()
		})
	}

	// setup client
	conn, err := grpc.NewClient(
		"static:///service",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(snappy.Name)),
	)
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		conn.Close()
	})
	s.client = pkg.NewServiceClient(conn)
}

func (s *GrpcServiceSuite) TestGet() {
	s.T().Parallel()

	ctx := context.Background()
	req := &pkg.GetRequest{Id: gofakeit.Int64()}
	resp, err := s.client.Get(ctx, req, grpc.UseCompressor(gzip.Name))
	s.Require().NoError(err)
	s.Require().Equal(req.GetId(), resp.GetId())
	s.Require().NotEmpty(resp.GetText())
}

func (s *GrpcServiceSuite) TestDownload() {
	s.T().Parallel()

	ctx := context.Background()
	req := &pkg.DownloadRequest{Width: 100, Height: 100}
	stream, err := s.client.Download(ctx, req)
	s.Require().NoError(err)

	var content []byte
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		s.Require().NoError(err)
		content = append(content, resp.GetChunk()...)
	}
	s.Require().NotEmpty(content)
	_, format, err := image.Decode(bytes.NewReader(content))
	s.Require().NoError(err)
	s.Require().Equal("jpeg", format)
}

func (s *GrpcServiceSuite) TestUpload() {
	s.T().Parallel()

	ctx := context.Background()
	stream, err := s.client.Upload(ctx)
	s.Require().NoError(err)

	content := gofakeit.ImageJpeg(100, 100)
	const chunkSize = 512

	for i := 0; i < len(content); i += chunkSize {
		err := stream.Send(&pkg.UploadRequest{Chunk: content[i:min(i+chunkSize, len(content))]})
		s.Require().NoError(err)
	}
	resp, err := stream.CloseAndRecv()
	s.Require().NoError(err)
	s.Require().Equal(int64(len(content)), resp.GetSize())
}

func Test(t *testing.T) {
	suite.Run(t, new(GrpcServiceSuite))
}
