package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Book struct {
	ID          int64           `db:"id"`
	PublisherID int64           `db:"publisher_id"`
	Title       string          `db:"title"`
	PublishDate time.Time       `db:"publish_date"`
	Price       decimal.Decimal `db:"price"`
	Pages       int64           `db:"pages"`
}
