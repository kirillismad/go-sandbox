package models

type Publisher struct {
	ID   int64  `db:"publishers.id"`
	Name string `db:"publishers.name"`
}

const (
	PublisherTable   = "publishers"
	PublisherColID   = "id"
	PublisherColName = "name"
)
