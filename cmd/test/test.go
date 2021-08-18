package main

import (
	"das-frama/zhukbot-tg/pkg/bot"
	"das-frama/zhukbot-tg/pkg/command"
	"das-frama/zhukbot-tg/pkg/config"
	"das-frama/zhukbot-tg/pkg/txtdb"
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

	cfg.TxtDB = &db

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
			result, err := command.Process(update.Message.Command(), update.Message, cfg)
			if err != nil {
				log.Printf("error: %v\n", err)
			}

			tbot.SendMessage(bot.SendMessageConfig{
				ChatID:              update.Message.Chat.ID,
				Text:                result.Text,
				DisableNotification: true,
			})
		}
	}

	log.Println("Shutting down...")
}
