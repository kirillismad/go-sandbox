package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"sandbox/utils"

	sb "github.com/huandu/go-sqlbuilder"
)

func (r *repo) CreatePubishers(ctx context.Context, items []entities.Publisher) ([]entities.Publisher, error) {
	query, args := r.createPublishersQuery(r.publisherToModelMany(items))

	var result []models.Publisher
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return nil, fmt.Errorf("r.db.SelectContext: %w", err)
	}

	return r.publisherToEntityMany(result), nil
}

func (r *repo) CreatePubisher(ctx context.Context, item entities.Publisher) (entities.Publisher, error) {
	result, err := r.CreatePubishers(ctx, []entities.Publisher{item})
	if err != nil {
		return entities.Publisher{}, fmt.Errorf("r.CreatePubishers: %w", err)
	}
	return result[0], nil
}

/* mapping */

func (r *repo) createPublishersQuery(items []models.Publisher) (sql string, args []interface{}) {
	b := sb.InsertInto(models.PublisherTable)
	b.Cols(models.PublisherColName)

	for _, v := range items {
		b.Values(v.Name)
	}

	b.SQL(returning(
		prfx(models.PublisherTable, models.PublisherColID),
		prfx(models.PublisherTable, models.PublisherColName),
	))

	query, args := b.Build()
	return query, args
}

func (r *repo) publisherToEntity(m models.Publisher) entities.Publisher {
	return entities.Publisher{
		ID:   m.ID,
		Name: m.Name,
	}
}

func (r *repo) publisherToEntityMany(slice []models.Publisher) []entities.Publisher {
	return utils.Map(slice, r.publisherToEntity)
}

func (r *repo) publisherToModel(e entities.Publisher) models.Publisher {
	return models.Publisher{
		ID:   e.ID,
		Name: e.Name,
	}
}

func (r *repo) publisherToModelMany(slice []entities.Publisher) []models.Publisher {
	return utils.Map(slice, r.publisherToModel)
}
