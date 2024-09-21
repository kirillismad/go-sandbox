package repo

import (
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repo_UpdateAuthor(t *testing.T) {
	t.Parallel()
	const query = `UPDATE authors SET name = $1 WHERE id = $2 RETURNING authors.id, authors.name`
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r, repoHandler, mock := initTest(t)

		q := mock.ExpectQuery(query)
		q.WithArgs()
		q.WillReturnRows(
			sqlmock.NewRows([]string{
				prfx(models.AuthorsTable, models.AuthorsColID),
				prfx(models.AuthorsTable, models.AuthorsColName),
			}).AddRow(
				1,
				"author1",
			),
		)

		result, err := repoHandler.GetRepo().UpdateAuthor(
			getctx(),
			entities.Author{
				ID:   1,
				Name: "author1",
			},
		)
		r.NoError(err)
		r.Equal(entities.Author{
			ID:   1,
			Name: "author1",
		}, result)
	})
}
