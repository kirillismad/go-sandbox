package repo

import (
	"context"
	// sb "github.com/huandu/go-sqlbuilder"
	"sandbox/sql/entities"
)

func (r *repo) UpsertAuthor(ctx context.Context, item entities.Author) (entities.Author, error) {
	return entities.Author{}, nil
}
