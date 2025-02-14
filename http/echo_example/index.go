package echo_example

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type Item struct {
	ID      int64                    `json:"id"`
	FInt    int64                    `json:"fInt"`
	FFloat  float64                  `json:"fFloat"`
	FString string                   `json:"fString"`
	FBool   bool                     `json:"fBool"`
	FSlice  []map[string]interface{} `json:"fSlice"`
}

var (
	storage = make(map[int64]*Item)
	mu      sync.RWMutex
)

func get(id int64) (*Item, bool) {
	mu.RLock()
	item, exists := storage[id]
	mu.RUnlock()
	return item, exists
}

func set(id int64, item *Item) {
	mu.Lock()
	storage[id] = item
	mu.Unlock()
}

func del(id int64) {
	mu.Lock()
	delete(storage, id)
	mu.Unlock()
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLogger(c echo.Context) *zap.Logger {
	return c.Get("logger").(*zap.Logger)
}

func getSingle(c echo.Context) error {
	logger := getLogger(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Failed to parse ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	item, exists := get(id)
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
	}
	return c.JSON(http.StatusOK, item)
}

func getList(c echo.Context) error {
	mu.RLock()
	items := lo.Values(storage)
	mu.RUnlock()
	return c.JSON(http.StatusOK, items)
}

func createItem(c echo.Context) error {
	logger := getLogger(c)
	var item Item
	if err := c.Bind(&item); err != nil {
		logger.Error("Failed to bind item", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	item.ID = gofakeit.Int64()
	set(item.ID, &item)
	return c.JSON(http.StatusCreated, item)
}

func deleteItem(c echo.Context) error {
	logger := getLogger(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Failed to parse ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	del(id)
	return c.NoContent(http.StatusNoContent)
}

func updateItem(c echo.Context) error {
	logger := getLogger(c)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Failed to parse ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	mu.Lock()
	existingItem, exists := storage[id]
	if !exists {
		mu.Unlock()
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
	}

	if err := c.Bind(existingItem); err != nil {
		logger.Error("Failed to bind item", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	mu.Unlock()
	return c.JSON(http.StatusOK, existingItem)
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
			logger.Info("Request",
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

	e.GET("/items/:id", getSingle)
	e.GET("/items", getList)
	e.POST("/items", createItem)
	e.DELETE("/items/:id", deleteItem)
	e.PUT("/items/:id", updateItem)

	return e
}
