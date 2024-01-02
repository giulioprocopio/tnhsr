package main

import (
	"fmt"
	"local/libs/dbutils"
	"os"
	"time"
)

func main() {
	conn := &dbutils.Conn{}

	conn.DSN = dbutils.DSN{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PSW"),
		Protocol: "tcp",
		Address:  os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_NAME"),
	}

	fmt.Printf("DSN is %s\n", conn.DSN.String())

	err := conn.Open()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Print("waiting for database...")
	conn.Wait(10)
	fmt.Println("\tdatabase available")

	version, err := conn.Version()
	if err != nil {
		panic(err)
	}
	fmt.Printf("using MySQL %s\n", version)

	iter := 0
	for {
		fmt.Printf("ping %d:\t", iter)

		err = conn.Ping()
		if err != nil {
			fmt.Printf("ping failed: %s", err.Error())
		} else {
			fmt.Println("OK")
		}

		// Delay 1 second.
		time.Sleep(time.Second)
		fmt.Printf("\033[1A\033[K") // Clear line
		iter++
	}
}
