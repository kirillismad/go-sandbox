package models

type Author struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type AuthorColumns struct {
	ID   string
	Name string
}

func (c AuthorColumns) All() []string {
	return []string{
		c.ID,
		c.Name,
	}
}

type TableMetadata[ColsT any] struct {
	TableName string
	Columns   ColsT
}

var AuthorMeta = TableMetadata[AuthorColumns]{
	TableName: "authors",
	Columns: AuthorColumns{
		ID:   "id",
		Name: "name",
	},
}
