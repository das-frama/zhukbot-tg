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
		ID:                      1,
		FirstName:               "Andrey",
		LastName:                "Galaktionov",
		Username:                "das_frama",
		LanguageCode:            "ru",
		CanJoinGroups:           true,
		CanReadAllGroupMessages: true,
	}

	db.Insert("users.txt", user)
}
