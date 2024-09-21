package repo

import (
	"context"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/mapper"
	"sandbox/sql/sqlx/models"
	"sandbox/utils"

	sb "github.com/huandu/go-sqlbuilder"
)

type authorCreator struct {
	mapper mapper.Mapper
}

func (b *authorCreator) MapToE(item models.Author) entities.Author {
	return b.mapper.AuthorToEntity(item)
}

func (b *authorCreator) MapToM(item entities.Author) models.Author {
	return b.mapper.AuthorToModel(item)
}

func (b *authorCreator) Query(in models.Author) (string, []interface{}) {
	builder := sb.InsertInto(models.AuthorsTable)
	builder.Cols(models.AuthorsColName)

	builder.Values(in.Name)

	builder.SQL(returning(
		prfx(models.AuthorsTable, models.AuthorsColID),
		prfx(models.AuthorsTable, models.AuthorsColName),
	))

	query, args := builder.Build()
	return query, args
}

func (r *repo) CreateAuthor(ctx context.Context, item entities.Author) (entities.Author, error) {
	t := NewCreator(
		&authorCreator{mapper: r.mapper},
		r.db,
	)
	return t.Create(ctx, item)
}

/* mapping */

func (r *repo) authorToEntity(m models.Author) entities.Author {
	return entities.Author{
		ID:   m.ID,
		Name: m.Name,
	}
}

func (r *repo) authorToEntityMany(slice []models.Author) []entities.Author {
	return utils.Map(slice, r.authorToEntity)
}

func (r *repo) authorToModel(e entities.Author) models.Author {
	return models.Author{
		ID:   e.ID,
		Name: e.Name,
	}
}

func (r *repo) authorToModelMany(slice []entities.Author) []models.Author {
	return utils.Map(slice, r.authorToModel)
}
