package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	Tx interface {
		Queryx(string, ...interface{}) (*sqlx.Rows, error)
		QueryRowx(string, ...interface{}) *sqlx.Row
		Exec(string, ...interface{}) (sql.Result, error)
		NamedExec(string, interface{}) (sql.Result, error)
		Select(interface{}, string, ...interface{}) error
		Get(interface{}, string, ...interface{}) error
		Rebind(string) string
		Commit() error
		Rollback() error
	}

	tx struct {
		client *sqlx.Tx
	}
)

func NewTx(client *sqlx.Tx) Tx {
	return &tx{client: client}
}

// Queryx queries the database and returns an *sqlx.Rows.
func (c *tx) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return c.client.Queryx(query, args...)
}

// QueryRowx queries the database and returns an *sqlx.Row.
func (c *tx) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return c.client.QueryRowx(query, args...)
}

// Exec using master db
func (c *tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.client.Exec(query, args...)
}

// Select using slave db.
func (c *tx) Select(dest interface{}, query string, args ...interface{}) error {
	return c.client.Select(dest, query, args...)
}

// Get using slave.
func (c *tx) Get(dest interface{}, query string, args ...interface{}) error {
	return c.client.Get(dest, query, args...)
}

// NamedExec using master db.
func (c *tx) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return c.client.NamedExec(query, arg)
}

// Rebind query
func (c *tx) Rebind(query string) string {
	return c.client.Rebind(query)
}

// Commit
func (c *tx) Commit() error {
	return c.client.Commit()
}

// Rollback
func (c *tx) Rollback() error {
	return c.client.Rollback()
}
