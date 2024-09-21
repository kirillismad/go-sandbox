package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"

	sb "github.com/huandu/go-sqlbuilder"
)

func (r *repo) UpdateAuthor(ctx context.Context, item entities.Author) (entities.Author, error) {
	m := r.authorToModel(item)

	b := sb.Update(models.AuthorsTable)
	b.Set(b.Assign(models.AuthorsColName, m.Name))
	b.Where(b.EQ(models.AuthorsColID, m.ID))
	b.SQL(returning(
		prfx(models.AuthorsTable, models.AuthorsColID),
		prfx(models.AuthorsTable, models.AuthorsColName),
	))

	query, args := b.Build()

	var result models.Author
	if err := r.db.GetContext(ctx, &result, query, args...); err != nil {
		return entities.Author{}, fmt.Errorf("r.db.GetContext: %w", err)
	}
	return r.authorToEntity(result), nil
}
