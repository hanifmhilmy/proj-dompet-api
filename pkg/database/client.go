package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	Client interface {
		SetMaxOpenConns(max int)
		SetMaxIdleConns(max int)
		Queryx(string, ...interface{}) (*sqlx.Rows, error)
		QueryRowx(string, ...interface{}) *sqlx.Row
		Exec(string, ...interface{}) (sql.Result, error)
		NamedExec(string, interface{}) (sql.Result, error)
		Select(interface{}, string, ...interface{}) error
		Get(interface{}, string, ...interface{}) error
		Beginx() (Tx, error)
		Rebind(string) string
		Ping() error
	}

	client struct {
		client *sqlx.DB
	}
)

func NewClient(db *sqlx.DB) Client {
	return &client{client: db}
}

// SetMaxOpenConnections to set max connections
func (c *client) SetMaxOpenConns(max int) {
	c.client.SetMaxOpenConns(max)
}

// SetMaxIdleConnections to set max idle connections
func (c *client) SetMaxIdleConns(max int) {
	c.client.SetMaxIdleConns(max)
}

// Queryx queries the database and returns an *sqlx.Rows.
func (c *client) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return c.client.Queryx(query, args...)
}

// QueryRowx queries the database and returns an *sqlx.Row.
func (c *client) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return c.client.QueryRowx(query, args...)
}

// Exec using master db
func (c *client) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.client.Exec(query, args...)
}

// Select using slave db.
func (c *client) Select(dest interface{}, query string, args ...interface{}) error {
	return c.client.Select(dest, query, args...)
}

// Get using slave.
func (c *client) Get(dest interface{}, query string, args ...interface{}) error {
	return c.client.Get(dest, query, args...)
}

// NamedExec using master db.
func (c *client) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return c.client.NamedExec(query, arg)
}

// Beginx sqlx transaction
func (c *client) Beginx() (Tx, error) {
	tx, err := c.client.Beginx()
	return NewTx(tx), err
}

// Rebind query
func (c *client) Rebind(query string) string {
	return c.client.Rebind(query)
}

//Ping ping db
func (c *client) Ping() error {
	return c.client.Ping()
}
