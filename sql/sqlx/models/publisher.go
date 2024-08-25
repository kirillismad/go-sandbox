package models

type Publisher struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
