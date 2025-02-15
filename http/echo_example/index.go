package echo_example

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type Product struct {
	ID       int64                    `json:"id"`
	Total    int64                    `json:"total"`
	Interest float64                  `json:"interest"`
	Title    string                   `json:"title"`
	IsActive bool                     `json:"is_active"`
	Content  []map[string]interface{} `json:"content"`
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLogger(c echo.Context) *zap.Logger {
	return c.Get("logger").(*zap.Logger)
}

func getHandler(c echo.Context) error {
	logger := getLogger(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Failed to parse ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	product := Product{
		ID:       id,
		Total:    gofakeit.Int64(),
		Interest: gofakeit.Float64(),
		Title:    gofakeit.Sentence(3),
		IsActive: gofakeit.Bool(),
		Content: []map[string]interface{}{
			{"key": gofakeit.Int8(), "value": gofakeit.Word()},
			{"key": gofakeit.Int8(), "value": gofakeit.Word()},
		},
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&product)
}

func listHandler(c echo.Context) error {
	products := lo.Times(5, func(i int) interface{} {
		return Product{
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
	})
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&products)
}

func createHandler(c echo.Context) error {
	logger := getLogger(c)

	var product Product
	if err := c.Bind(&product); err != nil {
		logger.Error("Failed to bind item", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	product.ID = gofakeit.Int64()
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusCreated)
	return json.NewEncoder(c.Response()).Encode(&product)
}

func deleteHandler(c echo.Context) error {
	logger := getLogger(c)
	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Failed to parse ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func updateHandler(c echo.Context) error {
	logger := getLogger(c)

	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Failed to parse ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var product Product
	if err := c.Bind(&product); err != nil {
		logger.Error("Failed to parse input", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&product)
}

func sendFile(c echo.Context) error {
	return c.File(os.Getenv("ECHO_EXAMPLE_FILE_PATH"))
}

func sendAttachment(c echo.Context) error {
	return c.Attachment(os.Getenv("ECHO_EXAMPLE_FILE_PATH"), "image.png")
}

func sendStream(c echo.Context) error {
	reader, writer := io.Pipe()
	defer reader.Close()

	go func() {
		defer writer.Close()
		for range 5 {
			p := Product{
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
			err := json.NewEncoder(writer).Encode(&p)
			if err != nil {
				getLogger(c).Error("Failed to encode product", zap.Error(err))
				return
			}
			_, err = writer.Write([]byte("\n"))
			if err != nil {
				getLogger(c).Error("Failed to add new line", zap.Error(err))
				return
			}
		}
	}()
	return c.Stream(http.StatusOK, echo.MIMEApplicationJSON, reader)
}

func BuildServer() *echo.Echo {
	logger, err := zap.NewProduction()
	fatalIfErr(err)

	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l := getLogger(c)
			l.Info("Request",
				zap.String("uri", v.URI),
				zap.String("method", v.Method),
				zap.Int("status", v.Status),
				zap.String("latency", v.Latency.String()),
			)
			return nil
		},
	}))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("logger", logger)
			return next(c)
		}
	})

	e.GET("/products/:id", getHandler)
	e.GET("/products", listHandler)
	e.POST("/products", createHandler)
	e.DELETE("/products/:id", deleteHandler)
	e.PUT("/products/:id", updateHandler)
	e.GET("/file", sendFile)
	e.GET("/attachment", sendAttachment)
	e.GET("/stream", sendStream)

	return e
}
