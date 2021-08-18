package main

import (
	"das-frama/zhukbot-tg/pkg/txtdb"
	"log"
)

func main() {
	db, err := txtdb.New("db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	user := txtdb.User{
		Username:     "Who_knowsme",
		FirstName:    "Sash",
		LastName:     "Nowikowa",
		LanguageCode: "ru",
	}

	db.Insert("users.txt", user)
	// db.Update("users.txt", user)
	// db.Delete("users.txt", user)
}
