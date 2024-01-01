package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DSN struct {
	Username, Password, Protocol, Address, Database string
	Options                                         map[string]string
}

func (dsn DSN) String() string {
	s := dsn.Username +
		":" + dsn.Password +
		"@" + dsn.Protocol +
		"(" + dsn.Address +
		")/" + dsn.Database

	if dsn.Options != nil {
		s += "?"
		for k, v := range dsn.Options {
			s += k + "=" + v + "&"
		}
		s = s[:len(s)-1]
	}

	return s
}

type DB struct {
	db  *sql.DB
	DSN DSN
}

func (db *DB) Open() error {
	var err error
	db.db, err = sql.Open("mysql", db.DSN.String())
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() error {
	return db.db.Close()
}
