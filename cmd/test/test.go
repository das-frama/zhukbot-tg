package main

import (
	"das-frama/zhukbot-tg/pkg/bot"
	"das-frama/zhukbot-tg/pkg/config"
	"das-frama/zhukbot-tg/pkg/txtdb"
	"das-frama/zhukbot-tg/pkg/zhuk"
	"log"
)

func main() {
	// Config.
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// DB.
	db, err := txtdb.New("db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Create telegram bot object.
	tbot := bot.New(cfg.Bot.Token)

	// Get updates channel.
	updates, err := tbot.GetUpdatesChan(bot.UpdateConfig{
		Offset:  0,
		Limit:   100,
		Timeout: 60,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Get through channels.
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Log incoming message.
		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)

		// Check if message is command.
		if update.Message.IsCommand() {
			// If command is /start.
			var err error
			switch update.Message.Command() {
			case "start":
				err = db.Insert("users.txt", txtdb.User{
					ID:                      update.Message.From.ID,
					Username:                update.Message.From.Username,
					FirstName:               update.Message.From.FirstName,
					LastName:                update.Message.From.LastName,
					LanguageCode:            update.Message.From.LanguageCode,
					CanJoinGroups:           update.Message.From.CanJoinGroups,
					CanReadAllGroupMessages: update.Message.From.CanReadAllGroupMessages,
				})
				if err != nil {
					log.Println(err)
				}

				err = db.Insert("chats.txt", txtdb.Chat{
					ID:            update.Message.Chat.ID,
					Type:          update.Message.Chat.Type,
					Title:         update.Message.Chat.Title,
					Username:      update.Message.Chat.Username,
					FirstName:     update.Message.Chat.FirstName,
					LastName:      update.Message.Chat.LastName,
					SlowModeDelay: update.Message.Chat.SlowModeDelay,
				})
				if err != nil {
					log.Println(err)
				}
			}
			// case "zhuk":
			// 	id, _ := db.CreateZhuk(postgres.Zhuk{
			// 		Name:   "Новый Жучара",
			// 		UserID: update.Message.From.ID,
			// 		ChatID: update.Message.Chat.ID,
			// 	})
			// 	_, err = tgBot.SendMessage(bot.SendMessageConfig{
			// 		ChatID: update.Message.Chat.ID,
			// 		Text:   fmt.Sprintf("Новый жучара с ID %d создан", id),
			// 	})
			// }
		} else {
			response, err := zhuk.AnalyzeText(update.Message.Text)
			if err != nil {
				log.Println(err)
			}
			if response != "" {
				_, err = tbot.SendMessage(bot.SendMessageConfig{
					ChatID: update.Message.Chat.ID,
					Text:   response,
				})
				if err != nil {
					log.Println(err)
				}
			}
		}

	}

	log.Println("Shutting down...")
}
