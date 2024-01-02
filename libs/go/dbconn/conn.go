package dbconn

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Conn struct {
	DB      *sql.DB
	DSN     DSN
	Timeout time.Duration
}

func NewConn() *Conn {
	conn := &Conn{
		DSN: DSN{
			Options: make(map[string]string),
		},
		Timeout: 5 * time.Second,
	}

	return conn
}

// Get timeout context and cancel function for database operations.
func (conn *Conn) Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), conn.Timeout)
}

// Open database connection.
func (conn *Conn) Open() error {
	var err error

	str, err := conn.DSN.String()
	if err != nil {
		return err
	}

	conn.DB, err = sql.Open("mysql", str)
	if err != nil {
		return err
	}

	return nil
}

// Ping database.
func (conn *Conn) Ping() error {
	ctx, cancel := conn.Context()
	defer cancel()

	err := conn.DB.PingContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Check if database is available.
func (conn *Conn) IsUp() (bool, error) {
	err := conn.Ping()
	return err == nil, err
}

// Ping until database it is available.
func (conn *Conn) Wait(timeout time.Duration) error {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	var err error
	for {
		select {
		case <-t.C:
			err = conn.Ping()
			if err == nil {
				return nil
			}
		case <-time.After(timeout):
			return err
		}
	}
}

// Get database version string.
func (conn *Conn) Version() (string, error) {
	ctx, cancel := conn.Context()
	defer cancel()

	var version string
	err := conn.DB.QueryRowContext(ctx, "SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	return version, nil
}

// Close database connection.  Preferably use `defer conn.Close()` after
// `conn.Open()`.
func (conn *Conn) Close() error {
	return conn.DB.Close()
}
