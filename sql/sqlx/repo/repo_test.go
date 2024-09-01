package repo

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func initTest(t *testing.T) (*require.Assertions, RepoHandler, sqlmock.Sqlmock) {
	r := require.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	r.NoError(err)

	repoHandler := NewRepoHandler(sqlx.NewDb(db, "pgx"))

	return r, repoHandler, mock
}

func getctx() context.Context {
	return context.Background()
}
