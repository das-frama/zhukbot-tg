package main

import (
	"log"
)

func main() {
	// Config.
	// cfg, err := config.LoadConfig("config.json")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Setup dependencies.

	// // DB Struct.

	// // Create telegram bot object.
	// tgBot := bot.New(cfg.Bot.Token)

	// // Get updates channel.
	// updates, err := tgBot.GetUpdatesChan(bot.UpdateConfig{
	// 	Offset:  0,
	// 	Limit:   100,
	// 	Timeout: 60,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Get through channels.
	// for update := range updates {
	// 	if update.Message == nil {
	// 		continue
	// 	}

	// 	// Log incoming message.
	// 	log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)

	// 	// Check if message is command.
	// 	if update.Message.IsCommand() {
	// 		// If command is /start.
	// 		var err error
	// 		switch update.Message.Command() {
	// 		case "start":
	// 			err = db.CreateUser(postgres.User{
	// 				ID:                      update.Message.From.ID,
	// 				FirstName:               update.Message.From.FirstName,
	// 				LastName:                update.Message.From.LastName,
	// 				Username:                update.Message.From.Username,
	// 				LanguageCode:            update.Message.From.LanguageCode,
	// 				CanJoinGroups:           update.Message.From.CanJoinGroups,
	// 				CanReadAllGroupMessages: update.Message.From.CanReadAllGroupMessages,
	// 			})
	// 			err = db.CreateChat(postgres.Chat{
	// 				ID:            update.Message.Chat.ID,
	// 				Type:          update.Message.Chat.Type,
	// 				Title:         update.Message.Chat.Title,
	// 				Username:      update.Message.Chat.Username,
	// 				FirstName:     update.Message.Chat.FirstName,
	// 				LastName:      update.Message.Chat.LastName,
	// 				SlowModeDelay: update.Message.Chat.SlowModeDelay,
	// 			})
	// 		case "zhuk":
	// 			id, _ := db.CreateZhuk(postgres.Zhuk{
	// 				Name:   "Новый Жучара",
	// 				UserID: update.Message.From.ID,
	// 				ChatID: update.Message.Chat.ID,
	// 			})
	// 			_, err = tgBot.SendMessage(bot.SendMessageConfig{
	// 				ChatID: update.Message.Chat.ID,
	// 				Text:   fmt.Sprintf("Новый жучара с ID %d создан", id),
	// 			})
	// 		}

	// 		if err != nil {
	// 			log.Println(err.Error())
	// 		}
	// 	} else {
	// 		response, err := zhuk.AnalyzeText(update.Message.Text)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 		}
	// 		if response != "" {
	// 			_, err = tgBot.SendMessage(bot.SendMessageConfig{
	// 				ChatID: update.Message.Chat.ID,
	// 				Text:   response,
	// 			})
	// 		}
	// 	}

	// }

	log.Println("Shutting down...")
}
