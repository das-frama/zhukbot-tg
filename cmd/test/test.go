package main

import (
	"das-frama/zhukbot-tg/pkg/txtdb"
	"fmt"
	"log"
)

func main() {
	db, err := txtdb.New("db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	fmt.Println(db)
}
