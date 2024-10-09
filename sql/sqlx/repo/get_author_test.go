package repo

import (
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repo_GetAuthor(t *testing.T) {
	t.Parallel()

	const query = `SELECT authors.id, authors.name FROM authors WHERE authors.id = $1`

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		r, repoHandler, mock := initTest(t)

		q := mock.ExpectQuery(query)
		q.WithArgs(1)
		q.WillReturnRows(sqlmock.NewRows([]string{
			prefix(models.AuthorsTable, models.AuthorsColID),
			prefix(models.AuthorsTable, models.AuthorsColName),
		}).AddRow(
			1,
			"author1",
		))

		result, err := repoHandler.GetRepo().GetAuthor(getctx(), GetAuthorParams{ID: 1})
		r.NoError(err)
		r.Equal(
			entities.Author{
				ID:   1,
				Name: "author1",
			},
			result,
		)
	})
}
