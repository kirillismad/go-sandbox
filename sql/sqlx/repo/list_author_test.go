package repo

import (
	"context"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repo_ListAuthor(t *testing.T) {
	t.Parallel()

	const query = `SELECT authors.id, authors.name FROM authors WHERE authors.name LIKE $1 LIMIT 5`

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		r, repoHandler, mock := initTest(t)

		q1 := mock.ExpectQuery(query)
		q1.WithArgs("author")
		q1.WillReturnRows(
			sqlmock.NewRows([]string{
				prfx(models.AuthorsTable, models.AuthorsColID),
				prfx(models.AuthorsTable, models.AuthorsColName),
			}).AddRow(11, "author1").AddRow(12, "author2"),
		)

		result, err := repoHandler.GetRepo().ListAuthor(context.Background(), ListAuthorParams{
			Offset:   0,
			Limit:    5,
			NameLike: "author",
		})

		r.NoError(err)
		r.Equal(
			[]entities.Author{
				{
					ID:   11,
					Name: "author1",
				},
				{
					ID:   12,
					Name: "author2",
				},
			},
			result,
		)
	})
}
