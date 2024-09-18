package repo

import (
	"context"
	"fmt"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"sandbox/utils"

	sb "github.com/huandu/go-sqlbuilder"
)

func (r *repo) ListAuthor(ctx context.Context) ([]entities.Author, error) {
	b := sb.Select(
		prfx(models.AuthorsTable, models.AuthorsColID),
		prfx(models.AuthorsTable, models.AuthorsColName),
	)
	b.From(models.AuthorsTable)

	query, args := b.Build()

	var result []models.Author
	if err := r.db.SelectContext(ctx, &result, query, args...); err != nil {
		return nil, fmt.Errorf("r.db.SelectContext: %w", err)
	}

	idList := utils.Map(result, func(item models.Author) interface{} {
		return item.ID
	})

	b = sb.Select(
		prfx(models.BookAuthorTable, models.BookAuthorColAuthorID),
		prfx(models.BookTable, models.BookColID),
		prfx(models.BookTable, models.BookColPublisherID),
		prfx(models.BookTable, models.BookColTitle),
		prfx(models.BookTable, models.BookColPublishDate),
		prfx(models.BookTable, models.BookColPrice),
		prfx(models.BookTable, models.BookColPages),
		prfx(models.PublisherTable, models.PublisherColID),
		prfx(models.PublisherTable, models.PublisherColName),
	)
	b.From(models.BookTable)
	b.JoinWithOption(
		sb.InnerJoin,
		models.BookAuthorTable,
		fmt.Sprintf(
			"%s = %s",
			prfx(models.BookAuthorTable, models.BookAuthorColBookID),
			prfx(models.BookTable, models.BookColID),
		),
	)
	b.JoinWithOption(
		sb.LeftJoin,
		models.PublisherTable,
		fmt.Sprintf(
			"%s = %s",
			prfx(models.BookTable, models.BookColPublisherID),
			prfx(models.PublisherTable, models.PublisherColID),
		),
	)
	b.Where(b.Any(
		prfx(models.BookAuthorTable, models.BookAuthorColID),
		"=",
		idList...,
	))

	query, args = b.Build()

	type helper struct {
		AuthorID int64 `db:"books_authors.author_id"`
		models.Book
		models.Publisher
	}

	var result2 []helper
	if err := r.db.SelectContext(ctx, &result2, query, args...); err != nil {
		return nil, fmt.Errorf("r.db.SelectContext: %w", err)
	}

	result2Map := make(map[int64][]helper)
	for _, item := range result2 {
		result2Map[item.AuthorID] = append(result2Map[item.AuthorID], item)
	}

	return utils.Map(result, func(item models.Author) entities.Author {
		r := entities.Author{
			ID:   item.ID,
			Name: item.Name,
		}

		sl, ok := result2Map[item.ID]
		if ok {
			for _, h := range sl {
				r.Books = append(r.Books, entities.Book{
					ID:          h.Book.ID,
					PublisherID: h.Book.PublisherID,
					Title:       h.Book.Title,
					PublishDate: h.Book.PublishDate,
					Price:       h.Book.Price,
					Pages:       h.Book.Pages,
					Publisher: entities.Publisher{
						ID:   h.Publisher.ID,
						Name: h.Publisher.Name,
					},
				})
			}
		}
		return r
	}), nil
}
