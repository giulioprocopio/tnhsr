package main

import (
	"fmt"
	"local/libs/auth/db"
	"os"
)

func main() {
	var tnhsrDB db.DB

	tnhsrDB.DSN = db.DSN{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PSW"),
		Protocol: "tcp",
		Address:  "localhost:3306",
		Database: os.Getenv("DB_NAME"),
	}

	fmt.Printf("DSN: %s\n", tnhsrDB.DSN.String())
	err := tnhsrDB.Open()
	if err != nil {
		panic(err)
	}

	defer tnhsrDB.Close()

	for {
		// Do nothing.
	}
}
