package dbconn

import "fmt"

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
