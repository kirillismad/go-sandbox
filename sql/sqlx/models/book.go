package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Book struct {
	ID          int64           `db:"books.id"`
	PublisherID int64           `db:"books.publisher_id"`
	Title       string          `db:"books.title"`
	PublishDate time.Time       `db:"books.publish_date"`
	Price       decimal.Decimal `db:"books.price"`
	Pages       int64           `db:"books.pages"`
}

const (
	BookTable          = "books"
	BookColID          = "id"
	BookColPublisherID = "publisher_id"
	BookColTitle       = "title"
	BookColPublishDate = "publish_date"
	BookColPrice       = "price"
	BookColPages       = "pages"
)
