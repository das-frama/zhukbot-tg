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

	user, err := db.FetchByUsername("users.txt", "Who_knowsme")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)

	err = db.Delete("users.txt", user)
	if err != nil {
		log.Fatalln(err)
	}
}
