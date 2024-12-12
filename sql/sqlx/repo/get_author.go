package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"

	sb "github.com/huandu/go-sqlbuilder"
	"github.com/samber/lo"
)

type GetAuthorParams struct {
	ID int64
}

func (r *repo) GetAuthor(ctx context.Context, params GetAuthorParams) (entities.Author, error) {
	prfx := lo.Partial(prefix, models.AuthorsTable)

	b := sb.Select(
		prfx(models.AuthorsColID),
		prfx(models.AuthorsColName),
	)
	b.From(models.AuthorsTable)

	if params.ID != 0 {
		b.Where(b.EQ(prfx(models.AuthorsColID), params.ID))
	}

	query, args := b.Build()

	var result models.Author
	if err := r.db.GetContext(ctx, &result, query, args...); err != nil {
		return entities.Author{}, fmt.Errorf("r.db.GetContext: %w", err)
	}

	return r.mapper.AuthorToEntity(result), nil
}
