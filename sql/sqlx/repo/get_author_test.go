package repo

import (
	"context"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
)

// "SELECT books_authors.author_id, books.id, books.publisher_id, books.title, books.publish_date, books.price, books.pages, publishers.id, publishers.name FROM books INNER JOIN books_authors ON books_authors.book_id = books.id LEFT JOIN publishers ON books.publisher_id = publishers.id WHERE books_authors.id = ANY ($1)"

func Test_repo_ListAuthor(t *testing.T) {
	t.Parallel()

	const query1 = `SELECT authors.id, authors.name FROM authors`
	const query2 = `
	SELECT books_authors.author_id, books.id, books.publisher_id, books.title, books.publish_date, books.price, books.pages, publishers.id, publishers.name 
	FROM books 
	INNER JOIN books_authors ON books_authors.book_id = books.id 
	LEFT JOIN publishers ON books.publisher_id = publishers.id 
	WHERE books_authors.id IN ($1, $2)
	`

	t.Run("success", func(t *testing.T) {
		r, repoHandler, mock := initTest(t)

		now := time.Now()

		q1 := mock.ExpectQuery(query1)
		q1.WithoutArgs()
		q1.WillReturnRows(
			sqlmock.NewRows([]string{
				prfx(models.AuthorsTable, models.AuthorsColID),
				prfx(models.AuthorsTable, models.AuthorsColName),
			}).AddRow(11, "author1").AddRow(12, "author2"),
		)

		q2 := mock.ExpectQuery(query2)
		q2.WithArgs(11, 12)
		q2.WillReturnRows(
			sqlmock.NewRows([]string{
				prfx(models.BookAuthorTable, models.BookAuthorColAuthorID),
				prfx(models.BookTable, models.BookColID),
				prfx(models.BookTable, models.BookColPublisherID),
				prfx(models.BookTable, models.BookColTitle),
				prfx(models.BookTable, models.BookColPublishDate),
				prfx(models.BookTable, models.BookColPrice),
				prfx(models.BookTable, models.BookColPages),
				prfx(models.PublisherTable, models.PublisherColID),
				prfx(models.PublisherTable, models.PublisherColName),
			}).AddRow(
				// books_authors.author_id
				11,
				// books.id
				22,
				// books.publisher_id
				33,
				// books.title
				"title1",
				// books.publish_date
				now,
				// books.price
				1_000,
				// books.pages
				300,
				// publishers.id
				33,
				// publishers.name
				"publisher1",
			).AddRow(
				// books_authors.author_id
				11,
				// books.id
				23,
				// books.publisher_id
				34,
				// books.title
				"title2",
				// books.publish_date
				now,
				// books.price
				1_000,
				// books.pages
				300,
				// publishers.id
				34,
				// publishers.name
				"publisher2",
			),
		)

		result, err := repoHandler.GetRepo().ListAuthor(context.Background())

		r.NoError(err)
		r.Equal([]entities.Author{
			{
				ID:   11,
				Name: "author1",
				Books: []entities.Book{
					{
						ID:          22,
						PublisherID: 33,
						Title:       "title1",
						PublishDate: now,
						Price:       decimal.NewFromInt(1000),
						Pages:       300,
						Publisher: entities.Publisher{
							ID:   33,
							Name: "publisher1",
						},
					},
					{
						ID:          23,
						PublisherID: 34,
						Title:       "title2",
						PublishDate: now,
						Price:       decimal.NewFromInt(1000),
						Pages:       300,
						Publisher: entities.Publisher{
							ID:   34,
							Name: "publisher2",
						},
					},
				},
			},
			{
				ID:   12,
				Name: "author2",
			},
		}, result)
	})
}
