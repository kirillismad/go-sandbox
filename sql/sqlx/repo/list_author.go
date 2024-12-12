package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"sandbox/utils"

	sb "github.com/huandu/go-sqlbuilder"
	"github.com/samber/lo"
)

type ListAuthorParams struct {
	Offset   int
	Limit    int
	NameLike string
}

func (r *repo) ListAuthor(ctx context.Context, params ListAuthorParams) ([]entities.Author, error) {
	prfx := lo.Partial(prefix, models.AuthorsTable)

	b := sb.Select(
		prfx(models.AuthorsColID),
		prfx(models.AuthorsColName),
	)
	b.From(models.AuthorsTable)

	if params.NameLike != "" {
		b.Where(b.Like(prfx(models.AuthorsColName), params.NameLike))
	}

	if params.Limit != 0 {
		b.Limit(params.Limit)
	}

	if params.Offset != 0 {
		b.Offset(params.Offset)
	}

	query, args := b.Build()

	var result []models.Author
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return nil, fmt.Errorf("r.db.SelectContext: %w", err)
	}

	return utils.Map(result, r.mapper.AuthorToEntity), nil
}
