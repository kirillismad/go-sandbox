package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"

	sb "github.com/huandu/go-sqlbuilder"
)

func (r *repo) DeleteAuthor(ctx context.Context, item entities.Author) error {
	m := r.authorToModel(item)

	b := sb.DeleteFrom(models.AuthorsTable)
	b.Where(b.EQ(models.AuthorsColID, m.ID))

	query, args := b.Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("r.db.ExecContext: %w", err)
	}

	return nil
}
