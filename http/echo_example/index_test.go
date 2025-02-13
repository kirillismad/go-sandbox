package echo_example

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (s *EchoSuite) SetupSuite() {
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
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			IdleConnTimeout: 20 * time.Second,
		},
		Timeout: 10 * time.Second,
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

func (s *EchoSuite) createTestItem() Item {
	item := Item{
		FInt:    gofakeit.Int64(),
		FFloat:  gofakeit.Float64(),
		FString: gofakeit.Word(),
		FBool:   gofakeit.Bool(),
		FSlice:  []map[string]interface{}{{"key": gofakeit.Int64(), "value": gofakeit.Word()}},
	}
	body, err := json.Marshal(item)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPost, "https://localhost:8080/items", bytes.NewBuffer(body))
	s.Require().NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	s.Require().NoError(err)

	err = json.Unmarshal(body, &item)
	s.Require().NoError(err)

	s.Equal(http.StatusCreated, resp.StatusCode)
	return item
}

func (s *EchoSuite) TestCreateItem() {
	s.createTestItem()
}

func (s *EchoSuite) TestGetList() {
	req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/items", nil)
	s.Require().NoError(err)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestDeleteItem() {
	item := s.createTestItem()

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("https://localhost:8080/items/%d", item.ID), nil)
	s.Require().NoError(err)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusNoContent, resp.StatusCode)
}

func (s *EchoSuite) TestUpdateItem() {
	item := s.createTestItem()

	item.FString = gofakeit.Word()
	body, err := json.Marshal(item)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("https://localhost:8080/items/%d", item.ID), bytes.NewBuffer(body))
	s.Require().NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestGetSingle() {
	item := s.createTestItem()

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://localhost:8080/items/%d", item.ID), nil)
	s.Require().NoError(err)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *EchoSuite) TestGetSingleNotFound() {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://localhost:8080/items/%d", gofakeit.Int64()), nil)
	s.Require().NoError(err)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *EchoSuite) TestUpdateItemNotFound() {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("https://localhost:8080/items/%d", gofakeit.Int64()), nil)
	s.Require().NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusNotFound, resp.StatusCode)
}

func TestEchoSuite(t *testing.T) {
	suite.Run(t, new(EchoSuite))
}
