package concrete

import (
	"context"
	"database/sql"
)

// DbConn is the concrete implementation of sqlGdbc by using *sql.DB
type DbConn struct {
	DB *sql.DB
}

// TxConn is the concrete implementation of sqlGdbc by using *sql.Tx
type TxConn struct {
	DB *sql.Tx
}

func (dbconn *DbConn) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return dbconn.DB.ExecContext(ctx, query, args...)
}

func (dbconn *DbConn) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return dbconn.DB.PrepareContext(ctx, query)
}

func (dbconn *DbConn) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return dbconn.DB.QueryContext(ctx, query, args...)
}

func (dbconn *DbConn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return dbconn.DB.QueryRowContext(ctx, query, args...)
}

func (txconn *TxConn) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return txconn.DB.ExecContext(ctx, query, args...)
}

func (txconn *TxConn) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return txconn.DB.PrepareContext(ctx, query)
}

func (txconn *TxConn) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return txconn.DB.QueryContext(ctx, query, args...)
}

func (txconn *TxConn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return txconn.DB.QueryRowContext(ctx, query, args...)
}
