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

	user1 := txtdb.User{
		Username:     "Who_knowsme",
		FirstName:    "Sash",
		LastName:     "Nowikowa",
		LanguageCode: "ru",
	}
	user2 := txtdb.User{
		Username:     "das_frama",
		FirstName:    "Andrey",
		LastName:     "Galaktionov",
		LanguageCode: "ru",
	}
	user3 := txtdb.User{
		Username:     "nickmann",
		FirstName:    "Nikita",
		LastName:     "Makarow",
		LanguageCode: "ru",
	}
	user4 := txtdb.User{
		Username:     "test",
		FirstName:    "Test",
		LastName:     "Test",
		LanguageCode: "ru",
	}

	db.Insert("users.txt", user1)
	db.Insert("users.txt", user2)
	db.Insert("users.txt", user4)
	db.Insert("users.txt", user3)

	user2.LanguageCode = "en"
	db.Update("users.txt", user2)

	db.Delete("users.txt", user4)
}
