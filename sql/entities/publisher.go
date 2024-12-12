package entities

type Publisher struct {
	ID   int64
	Name string

	Books []Book
}
