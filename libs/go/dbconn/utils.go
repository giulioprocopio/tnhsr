package dbconn

import (
	"errors"
	"os"
)

func (conn *Conn) ExecFileUnsafe(path string) (err error) {
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
