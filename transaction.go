package zenorm

import "database/sql"
import "context"

// compiler checker
var _ Queryer = (*transaction)(nil)

type transaction struct {
	*sql.Tx
}

func (tx *transaction) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return tx.QueryContext(context.TODO(), query, args...)
}

func (tx *transaction) QueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return parseRows(tx.Tx.QueryContext(ctx, query, args...))
}

func (tx *transaction) QueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	return tx.QueryRowContext(context.TODO(), query, args...)
}

func (tx *transaction) QueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	return parseRow(tx.Tx.QueryRowContext(ctx, query, args...))
}

func (tx *transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.Tx.Exec(query, args)
}