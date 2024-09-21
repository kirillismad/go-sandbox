package repo

import (
	"math/rand/v2"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/lo"
)

func TestCreateAuthors(t *testing.T) {
	t.Parallel()

	const query = `INSERT INTO authors (name) VALUES ($1) RETURNING authors.id, authors.name`
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
				prfx(models.AuthorsTable, models.AuthorsColID),
				prfx(models.AuthorsTable, models.AuthorsColName),
			}).AddRow(ID, entity.Name),
		)

		//act
		result, err := repoHandler.GetRepo().CreateAuthor(getctx(), entity)

		// assert
		r.NoError(err)
		r.Equal(entities.Author{ID: ID, Name: entity.Name}, result)
	})

	t.Run("success: in tx", func(t *testing.T) {
		t.Parallel()

		r, repoHandler, mock := initTest(t)

		// arrange
		ID, entity := genEntity()

		mock.ExpectBegin()
		q := mock.ExpectQuery(query)
		q.WithArgs(entity.Name)
		q.WillReturnRows(
			sqlmock.NewRows([]string{
				prfx(models.AuthorsTable, models.AuthorsColID),
				prfx(models.AuthorsTable, models.AuthorsColName),
			}).AddRow(ID, entity.Name),
		)
		mock.ExpectCommit()

		//act
		var result entities.Author
		err := repoHandler.InTrasaction(func(repo Repo) error {
			var errTx error
			result, errTx = repo.CreateAuthor(getctx(), entity)
			return errTx
		})

		// assert
		r.NoError(err)
		r.Equal(entities.Author{ID: ID, Name: entity.Name}, result)
	})
}
