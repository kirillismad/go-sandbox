package echo_example

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/http2"
)

type EchoSuite struct {
	suite.Suite
	server *echo.Echo
	ch     <-chan error
	client *http.Client
}

func TestEchoSuite(t *testing.T) {
	suite.Run(t, new(EchoSuite))
}

func (s *EchoSuite) SetupSuite() {
	path, err := filepath.Abs("./testdata/websocket.png")
	s.Require().NoError(err)

	s.T().Setenv("ECHO_EXAMPLE_FILE_PATH", path)

	s.server = BuildServer()

	s.ch = lo.Async(func() error {
		err := s.server.StartTLS(":8080", "./testdata/cert.pem", "./testdata/key.pem")
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	select {
	case err := <-s.ch:
		s.Require().NoError(err)
	case <-time.After(2 * time.Second):
	}
	s.client = &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			AllowHTTP:       true,
		},
	}
	s.T().Cleanup(func() {
		select {
		case err := <-s.ch:
			s.Require().NoError(err)
		default:
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.server.Shutdown(ctx)
		s.Require().NoError(err)
	})
}

func (s *EchoSuite) TestCreateProduct() {
	product := Product{
		ID:       gofakeit.Int64(),
		Total:    gofakeit.Int64(),
		Interest: gofakeit.Float64(),
		Title:    gofakeit.Sentence(3),
		IsActive: gofakeit.Bool(),
		Content: []map[string]interface{}{
			{"key": gofakeit.Int8(), "value": gofakeit.Word()},
			{"key": gofakeit.Int8(), "value": gofakeit.Word()},
		},
	}
	body, err := json.Marshal(product)
	s.Require().NoError(err)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/products", bytes.NewBuffer(body))
	s.Require().NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusCreated, resp.StatusCode)
}

func (s *EchoSuite) TestGetList() {
	req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/products", nil)
	s.Require().NoError(err)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)

	var products []Product
	err = json.Unmarshal(body, &products)
	s.Require().NoError(err)

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestDeleteItem() {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("https://localhost:8080/products/%d", gofakeit.Int64()), nil)
	s.Require().NoError(err)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusNoContent, resp.StatusCode)
}

func (s *EchoSuite) TestUpdateItem() {
	product := Product{
		ID:       gofakeit.Int64(),
		Total:    gofakeit.Int64(),
		Interest: gofakeit.Float64(),
		Title:    gofakeit.Sentence(3),
		IsActive: gofakeit.Bool(),
		Content: []map[string]interface{}{
			{"key": gofakeit.Int8(), "value": gofakeit.Word()},
			{"key": gofakeit.Int8(), "value": gofakeit.Word()},
		},
	}
	body, err := json.Marshal(product)
	s.Require().NoError(err)
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("https://localhost:8080/products/%d", product.ID), bytes.NewBuffer(body))
	s.Require().NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestGetSingle() {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://localhost:8080/products/%d", gofakeit.Int64()), nil)
	s.Require().NoError(err)
	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)

	var product Product
	err = json.Unmarshal(body, &product)
	s.Require().NoError(err)

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestDownloadFile() {
	req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/file", nil)
	s.Require().NoError(err)
	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.T().Logf("Header: %v", resp.Header)

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestAttachmentFile() {
	req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/attachment", nil)
	s.Require().NoError(err)
	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.T().Logf("Header: %v", resp.Header)

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestStream() {
	req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/stream", nil)
	s.Require().NoError(err)
	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.T().Logf("Header: %v", resp.Header)

	s.Equal(http.StatusOK, resp.StatusCode)
}
