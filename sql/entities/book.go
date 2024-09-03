package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type Book struct {
	ID          int64
	PublisherID int64
	Title       string
	PublishDate time.Time
	Price       decimal.Decimal
	Pages       int64

	Publisher Publisher
	Authors   []Author
}
