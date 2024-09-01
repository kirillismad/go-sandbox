package repo

import (
	"context"
	"database/sql"
	"fmt"
	"sandbox/sql/entities"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func init() {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
}

func returning(col ...string) string {
	return fmt.Sprintf("RETURNING %s", strings.Join(col, ", "))
}

type RepoHandler interface {
	GetRepo() Repo
	InTrasaction(funcTx func(repo Repo) error) error
}

type Repo interface {
	CreateAuthor(ctx context.Context, entity entities.Author) (entities.Author, error)
}

type repoHandler struct {
	db *sqlx.DB
}

func NewRepoHandler(db *sqlx.DB) RepoHandler {
	return &repoHandler{db: db}
}

func (h *repoHandler) GetRepo() Repo {
	return newRepo(h.db)
}

func (h *repoHandler) InTrasaction(funcTx func(r Repo) error) error {
	tx, err := h.db.Beginx()

	if err != nil {
		return fmt.Errorf("db.Beginx: %w", err)
	}
	defer tx.Rollback()

	if err := funcTx(newRepo(h.db)); err != nil {
		return fmt.Errorf("funcTx: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}

	return nil
}

type DBTX interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type repo struct {
	db DBTX
}

func newRepo(db DBTX) *repo {
	return &repo{db: db}
}
