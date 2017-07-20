package zenorm

import "database/sql"
import "context"

// compiler checker
var _ Queryer = (*transaction)(nil)

type transaction struct {
	*sql.Tx
}

func (tx *transaction) Query(query string, args ...interface{}) ([]map[string]string, error) {
	return tx.QueryContext(context.TODO(), query, args...)
}

func (tx *transaction) QueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]string, error) {
	return parseRows(tx.Tx.QueryContext(ctx, query, args...))
}

func (tx *transaction) QueryRow(query string, args ...interface{}) *sql.Row {
	return tx.QueryRowContext(context.TODO(), query, args...)
}

func (tx *transaction) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return tx.Tx.QueryRowContext(ctx, query, args...)
}

func (tx *transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.Tx.ExecContext(context.TODO(), query, args...)
}

func (tx *transaction) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.Tx.ExecContext(ctx, query, args...)
}
