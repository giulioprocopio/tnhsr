package dbconn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Refer to [go-sql-driver/mysql][1] for more information about params spec.
// [1]: https://github.com/go-sql-driver/mysql#parameters
type DSN struct {
	Username, Password, Protocol, Address, Database string
	Options                                         map[string]string
}

// Assert DSN is valid and ready to use.
func (dsn DSN) Ready() error {
	fields := map[string]string{
		"username":      dsn.Username,
		"password":      dsn.Password,
		"protocol":      dsn.Protocol,
		"address":       dsn.Address,
		"database name": dsn.Database,
	}

	for field, value := range fields {
		if value == "" {
			return fmt.Errorf("DSN %s is required", field)
		}
	}

	return nil
}

// Format DSN parameters to string:
// <username>:<password>@<protocol>(<address>)/<database>?<param>=<value>.
func (dsn DSN) String() (string, error) {
	s := dsn.Username +
		":" + dsn.Password +
		"@" + dsn.Protocol +
		"(" + dsn.Address +
		")/" + dsn.Database

	if dsn.Options != nil && len(dsn.Options) > 0 {
		s += "?"
		for k, v := range dsn.Options {
			s += k + "=" + v + "&"
		}
		s = s[:len(s)-1]
	}

	return s, dsn.Ready()
}

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

func (conn *Conn) ExecFileUnsafe(path string) error {
	if conn.DSN.Options == nil ||
		conn.DSN.Options["multiStatements"] != "true" {
		return errors.New(
			"DSN `multiStatements` option must be set to execute file")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	sql := string(content)

	ctx, cancel := conn.Context()
	defer cancel()

	_, err = conn.DB.ExecContext(ctx, sql)
	if err != nil {
		return err
	}

	return nil
}
