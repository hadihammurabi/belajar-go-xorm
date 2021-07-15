package main

import (
	"log"

	"belajar-go-xorm/connection"
	"belajar-go-xorm/insert"
	"belajar-go-xorm/query"
	"belajar-go-xorm/sync"
)

func main() {
	conn, err := connection.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connection defined")

	err = sync.Sync(conn)
	if err != nil {
		log.Fatalln(err)
	}

	affected, err := insert.Insert(conn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("inserted =>", affected)

	affecteds, err := insert.InsertMany(conn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("inserted =>", affecteds)

	getResult, err := query.Get(conn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("get result =>", getResult)

	findResult, err := query.Get(conn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("find result =>", findResult)

}
