package main

import (
	"fmt"
	"local/libs/auth/db"
	"os"
	"time"
)

func main() {
	tnhsrDB := &db.DB{}

	tnhsrDB.DSN = db.DSN{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PSW"),
		Protocol: "tcp",
		Address:  os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_NAME"),
	}

	fmt.Printf("DSN: %s\n", tnhsrDB.DSN.String())

	err := tnhsrDB.Open()
	if err != nil {
		panic(err)
	}
	defer tnhsrDB.Close()

	for {
		err = tnhsrDB.Ping()
		if err != nil {
			fmt.Printf("Ping failed: %s\n", err.Error())
		} else {
			fmt.Println("Ping OK")
		}

		// Delay 1 second.
		time.Sleep(5 * time.Second)
	}
}
