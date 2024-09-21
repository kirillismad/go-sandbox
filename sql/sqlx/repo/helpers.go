package repo

import (
	"context"
	"database/sql"
	"fmt"
)

type IGetContext interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type IExecContext interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type ICreator[M, E any] interface {
	Query(M) (string, []interface{})
	MapToM(E) M
	MapToE(M) E
}

type Creator[M, E any] struct {
	ICreator[M, E]
	db IGetContext
}

func NewCreator[M, E any](base ICreator[M, E], db IGetContext) *Creator[M, E] {
	return &Creator[M, E]{
		ICreator: base,
		db:       db,
	}
}

func (c Creator[M, E]) Create(ctx context.Context, item E) (E, error) {
	query, args := c.Query(c.MapToM(item))
	var result M
	if err := c.db.GetContext(ctx, &result, query, args...); err != nil {
		var zero E
		return zero, fmt.Errorf("GetContext: %w", err)
	}
	return c.MapToE(result), nil
}
