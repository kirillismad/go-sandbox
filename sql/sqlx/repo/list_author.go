package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"

	sb "github.com/huandu/go-sqlbuilder"
)

type ListAuthorParams struct {
	Offset   int
	Limit    int
	NameLike string
}

func (r *repo) ListAuthor(ctx context.Context, params ListAuthorParams) ([]entities.Author, error) {
	query, args := r.listAuthorQuery(params)

	var result []models.Author
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return nil, fmt.Errorf("r.db.SelectContext: %w", err)
	}

	return r.authorToEntityMany(result), nil
}

func (*repo) listAuthorQuery(params ListAuthorParams) (string, []interface{}) {
	b := sb.Select(
		prfx(models.AuthorsTable, models.AuthorsColID),
		prfx(models.AuthorsTable, models.AuthorsColName),
	)
	b.From(models.AuthorsTable)

	if params.NameLike != "" {
		b.Where(b.Like(prfx(models.AuthorsTable, models.AuthorsColName), params.NameLike))
	}

	if params.Limit != 0 {
		b.Limit(params.Limit)
	}

	if params.Offset != 0 {
		b.Offset(params.Offset)
	}

	query, args := b.Build()
	return query, args
}
