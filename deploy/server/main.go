package main

import (
	"fmt"
	"local/libs/auth/db"
	"os"
	"time"
)

func main() {
	tnhsr := &db.DB{}

	tnhsr.DSN = db.DSN{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PSW"),
		Protocol: "tcp",
		Address:  os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_NAME"),
	}

	fmt.Printf("DSN: %s\n", tnhsr.DSN.String())

	err := tnhsr.Open()
	if err != nil {
		panic(err)
	}
	defer tnhsr.Close()

	fmt.Println("Waiting for database...")
	tnhsr.Wait(10)
	fmt.Println("Database available")

	version, err := tnhsr.Version()
	if err != nil {
		panic(err)
	}
	fmt.Printf("MySQL Version: %s\n", version)

	for {
		err = tnhsr.Ping()
		if err != nil {
			fmt.Printf("Ping failed: %s\n", err.Error())
		} else {
			fmt.Println("Ping OK")
		}

		// Delay 1 second.
		time.Sleep(time.Second)
	}
}
