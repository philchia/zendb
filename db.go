package zenorm

import (
	"context"
	"database/sql"
)

// Database ...
type Database interface {
	Queryer
	Txer
	Closer
}

// Open a database
func Open(driverName string, dataSourceName string) (Database, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &database{db: db}, nil
}

// Queryer query data from database
type Queryer interface {
	Query(query string, args ...interface{}) ([]map[string]string, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]string, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// Closer close database
type Closer interface {
	Close() error
}

// Txer run execs on transaction
type Txer interface {
	Tx(f func(Queryer) error) error
}

type database struct {
	db *sql.DB
}

func (database *database) begin() (*transaction, error) {
	tx, err := database.db.Begin()
	if err != nil {
		return nil, err
	}
	return &transaction{
		Tx: tx,
	}, nil
}

func (database *database) Query(query string, args ...interface{}) ([]map[string]string, error) {
	return database.QueryContext(context.TODO(), query, args...)
}

func (database *database) QueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]string, error) {
	return parseRows(database.db.QueryContext(context.TODO(), query, args...))
}

func (database *database) QueryRow(query string, args ...interface{}) *sql.Row {
	return database.QueryRowContext(context.TODO(), query, args...)
}

func (database *database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return database.db.QueryRowContext(ctx, query, args...)
}

func (database *database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return database.db.Exec(query, args...)
}

// Tx wrap all db operations into a closure
func (database *database) Tx(f func(Queryer) error) (err error) {
	tx, err := database.begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return f(tx)
}

func (database *database) Close() error {
	return database.db.Close()
}
