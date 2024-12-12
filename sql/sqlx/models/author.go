package models

type Author struct {
	ID   int64  `db:"authors.id"`
	Name string `db:"authors.name"`
}

const (
	AuthorsTable   = "authors"
	AuthorsColID   = "id"
	AuthorsColName = "name"
)
