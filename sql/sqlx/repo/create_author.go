package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"sandbox/utils"

	sb "github.com/huandu/go-sqlbuilder"
)

func (r *repo) CreateAuthors(ctx context.Context, items []entities.Author) ([]entities.Author, error) {
	query, args := r.createAuthorsQuery(r.authorToModelMany(items))

	var result []models.Author
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return nil, fmt.Errorf("r.db.SelectContext: %w", err)
	}

	return r.authorToEntityMany(result), nil
}

func (r *repo) CreateAuthor(ctx context.Context, item entities.Author) (entities.Author, error) {
	result, err := r.CreateAuthors(ctx, []entities.Author{item})
	if err != nil {
		return entities.Author{}, fmt.Errorf("r.CreateAuthors: %w", err)
	}
	return result[0], nil
}

func (r *repo) createAuthorsQuery(items []models.Author) (sql string, args []interface{}) {
	b := sb.InsertInto(models.AuthorsTable)
	b.Cols(models.AuthorsColName)

	for _, v := range items {
		b.Values(v.Name)
	}

	b.SQL(returning(
		prfx(models.AuthorsTable, models.AuthorsColID),
		prfx(models.AuthorsTable, models.AuthorsColName),
	))

	query, args := b.Build()
	return query, args
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
