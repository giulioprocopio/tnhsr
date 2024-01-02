package main

import (
	"fmt"
	"local/libs/dbutils"
	"os"
	"time"
)

func main() {
	conn := dbutils.NewConn()

	conn.DSN.Username = os.Getenv("DB_USER")
	conn.DSN.Password = os.Getenv("DB_PSW")
	conn.DSN.Protocol = "tcp"
	conn.DSN.Address = os.Getenv("DB_ADDR")
	conn.DSN.Database = os.Getenv("DB_NAME")

	str, _ := conn.DSN.String()
	fmt.Printf("DSN is %s\n", str)

	err := conn.Open()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("waiting for database...")
	err = conn.Wait(time.Second * 5)
	if err != nil {
		panic(err)
	}
	fmt.Println("database is ready")

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
			fmt.Printf("ping failed: %s\n", err.Error())
		} else {
			fmt.Println("OK")
		}

		// Delay 1 second.
		time.Sleep(time.Second)
		// Clear line.
		fmt.Printf("\033[1A\033[K")
		iter++
	}
}
