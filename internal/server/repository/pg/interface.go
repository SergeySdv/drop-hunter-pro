package pg

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// this interface is a wrapper on sqlx.Tx and sqlx.DB
type SqlDriver interface {
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type Tx struct {
	origin *sqlx.Tx
}

func WrapSqlxTx(tx *sqlx.Tx) *Tx {
	return &Tx{
		origin: tx,
	}
}

func (tx *Tx) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	// todo: someday when https://github.com/jmoiron/sqlx/pull/565 will be merged all this file can be removed
	return tx.origin.NamedQuery(query, arg)
}

func (tx *Tx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return tx.origin.SelectContext(ctx, dest, query, args...)
}

func (tx *Tx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return tx.origin.GetContext(ctx, dest, query, args...)
}
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.origin.ExecContext(ctx, query, args...)
}

func (tx *Tx) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return tx.origin.NamedExecContext(ctx, query, arg)
}
func (tx *Tx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return tx.origin.QueryRowxContext(ctx, query, args...)
}
