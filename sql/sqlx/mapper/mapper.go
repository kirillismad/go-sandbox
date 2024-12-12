package mapper

import (
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
)

type Mapper interface {
	AuthorToModel(item entities.Author) models.Author
	AuthorToEntity(item models.Author) entities.Author
}

type mapper struct{}

func New() Mapper {
	return new(mapper)
}
