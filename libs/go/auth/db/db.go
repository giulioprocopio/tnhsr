package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DSN struct {
	Username, Password, Protocol, Address, Database string
	Options                                         map[string]string
}

// Format DSN parameters to string:
// <username>:<password>@<protocol>(<address>)/<database>?<param>=<value>.
func (dsn DSN) String() string {
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

	return s
}

type DB struct {
	DB  *sql.DB
	DSN DSN
}

// Open database connection.
func (db *DB) Open() error {
	var err error
	db.DB, err = sql.Open("mysql", db.DSN.String())
	if err != nil {
		return err
	}

	return nil
}

// Ping database.
func (db *DB) Ping() error {
	return db.DB.Ping()
}

// Ping until database is available.
func (db *DB) Wait(timeout int) error {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	var err error
	for {
		select {
		case <-t.C:
			err = db.Ping()
			if err == nil {
				return nil
			}
		case <-time.After(time.Duration(timeout) * time.Second):
			return err
		}
	}
}

// Get database version string.
func (db *DB) Version() (string, error) {
	var version string
	err := db.DB.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	return version, nil
}

// Close database connection.  Preferably use `defer db.Close()` after
// `db.Open()`.
func (db *DB) Close() error {
	return db.DB.Close()
}
