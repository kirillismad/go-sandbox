package repo

import (
	"context"
	"sandbox/sql/entities"
	"sandbox/sql/sqlx/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestCreateAuthor(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	r.NoError(err)

	q := mock.ExpectQuery(`INSERT INTO authors (name) VALUES ($1) RETURNING id, name`)
	q.WithArgs("John Doe")
	q.WillReturnRows(sqlmock.NewRows(models.AuthorMeta.Columns.All()).AddRow(1, "John Doe"))

	repo := newRepo(sqlx.NewDb(db, "pgx"))

	entity, err := repo.CreateAuthor(context.Background(), entities.Author{Name: "John Doe"})
	r.NoError(err)
	r.Equal(entities.Author{ID: 1, Name: "John Doe"}, entity)
}
