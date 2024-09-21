package repo

import (
	"sandbox/sql/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repo_DeleteAuthor(t *testing.T) {
	t.Parallel()
	const query = `DELETE FROM authors WHERE id = $1`
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r, repoHandler, mock := initTest(t)

		q := mock.ExpectExec(query)
		q.WithArgs(1)
		q.WillReturnResult(sqlmock.NewResult(0, 1))

		err := repoHandler.GetRepo().DeleteAuthor(
			getctx(),
			entities.Author{
				ID:   1,
				Name: "author1",
			},
		)
		r.NoError(err)
	})
}
