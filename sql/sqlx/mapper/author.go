package mapper

import (
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
)

func (m *mapper) AuthorToModel(item entities.Author) models.Author {
	return models.Author{
		ID:   item.ID,
		Name: item.Name,
	}
}

func (m *mapper) AuthorToEntity(item models.Author) entities.Author {
	return entities.Author{
		ID:   item.ID,
		Name: item.Name,
	}
}
