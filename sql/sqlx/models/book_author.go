package models

type BookAuthorColumns struct {
	ID       string `db:"books_authors.id"`
	AuthorID string `db:"books_authors.author_id"`
	BookID   string `db:"books_authors.book_id"`
}

const (
	BookAuthorTable       = "books_authors"
	BookAuthorColID       = "id"
	BookAuthorColAuthorID = "author_id"
	BookAuthorColBookID   = "book_id"
)
