package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

func (r *repo) CreateAuthor(ctx context.Context, entity entities.Author) (entities.Author, error) {
	b := sqlbuilder.InsertInto(models.AuthorMeta.TableName)
	b.Cols(models.AuthorMeta.Columns.Name)
	b.Values(entity.Name)
	b.SQL(fmt.Sprintf("RETURNING %s", strings.Join(models.AuthorMeta.Columns.All(), ", ")))

	query, args := b.Build()

	var model models.Author
	if err := r.db.GetContext(ctx, &model, query, args...); err != nil {
		return entities.Author{}, fmt.Errorf("r.db.GetContext: %w", err)
	}

	return entities.Author{
		ID:   model.ID,
		Name: model.Name,
	}, nil
}
