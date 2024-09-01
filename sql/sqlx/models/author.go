package models

type Author struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type AuthorColumns struct {
	ID   string
	Name string
}

var AuthorMeta = TableMetadata[AuthorColumns]{
	TableName: "authors",
	Columns: AuthorColumns{
		ID:   "id",
		Name: "name",
	},
}
