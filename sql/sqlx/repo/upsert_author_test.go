package repo

import (
	"math/rand/v2"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/samber/lo"
)

func Test_repo_UpsertAuthor(t *testing.T) {
	t.Parallel()

	const query = `INSERT INTO authors (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING authors.id, authors.name`
	genEntity := func() (int64, entities.Author) {
		entity := entities.Author{
			Name: lo.RandomString(8, lo.LettersCharset),
		}
		ID := rand.Int64()
		return ID, entity
	}

	t.Run("success: no tx", func(t *testing.T) {
		t.Parallel()

		r, repoHandler, mock := initTest(t)

		// arrange
		ID, entity := genEntity()

		q := mock.ExpectQuery(query)
		q.WithArgs(entity.Name)
		q.WillReturnRows(
			sqlmock.NewRows([]string{
				prefix(models.AuthorsTable, models.AuthorsColID),
				prefix(models.AuthorsTable, models.AuthorsColName),
			}).AddRow(ID, entity.Name),
		)

		//act
		result, err := repoHandler.GetRepo().UpsertAuthor(getctx(), entity)

		// assert
		r.NoError(err)
		r.Equal(entities.Author{ID: ID, Name: entity.Name}, result)
	})
}
