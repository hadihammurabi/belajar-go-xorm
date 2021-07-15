package main

import (
	"log"

	"belajar-go-xorm/connection"
	"belajar-go-xorm/sync"
)

func main() {
	conn, err := connection.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(conn)
}
