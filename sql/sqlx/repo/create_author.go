package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"

	sb "github.com/huandu/go-sqlbuilder"
	"github.com/samber/lo"
)

func (r *repo) CreateAuthor(ctx context.Context, item entities.Author) (entities.Author, error) {
	m := r.mapper.AuthorToModel(item)

	prfx := lo.Partial(prefix, models.AuthorsTable)

	b := sb.InsertInto(models.AuthorsTable)
	b.Cols(models.AuthorsColName)
	b.Values(m.Name)
	b.SQL(returning(
		prfx(models.AuthorsColID),
		prfx(models.AuthorsColName),
	))

	query, args := b.Build()

	var result models.Author
	if err := r.db.GetContext(ctx, &result, query, args...); err != nil {
		return entities.Author{}, fmt.Errorf("GetContext: %w", err)
	}
	return r.mapper.AuthorToEntity(result), nil
}
