package main

import (
	"context"
	"das-frama/zhukbot-tg/pkg/bot"
	"das-frama/zhukbot-tg/pkg/config"
	"das-frama/zhukbot-tg/pkg/postgres"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Config.
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Setup dependencies.
	// DB Conn.
	conn, err := pgxpool.Connect(context.Background(), cfg.DB.URL)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	// DB Struct.
	db := postgres.New(conn)

	// Create telegram bot object.
	tgBot := bot.New(cfg.Bot.Token)

	// Get updates channel.
	updates, err := tgBot.GetUpdatesChan(bot.UpdateConfig{
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
				err = db.CreateUser(postgres.User{
					ID:                      update.Message.From.ID,
					FirstName:               update.Message.From.FirstName,
					LastName:                update.Message.From.LastName,
					Username:                update.Message.From.Username,
					LanguageCode:            update.Message.From.LanguageCode,
					CanJoinGroups:           update.Message.From.CanJoinGroups,
					CanReadAllGroupMessages: update.Message.From.CanReadAllGroupMessages,
				})
				err = db.CreateChat(postgres.Chat{
					ID:            update.Message.Chat.ID,
					Type:          update.Message.Chat.Type,
					Title:         update.Message.Chat.Title,
					Username:      update.Message.Chat.Username,
					FirstName:     update.Message.Chat.FirstName,
					LastName:      update.Message.Chat.LastName,
					SlowModeDelay: update.Message.Chat.SlowModeDelay,
				})
			case "zhuk":
				id, _ := db.CreateZhuk(postgres.Zhuk{
					Name:   "Новый Жучара",
					UserID: update.Message.From.ID,
					ChatID: update.Message.Chat.ID,
				})
				_, err = tgBot.SendMessage(bot.SendMessageConfig{
					ChatID: update.Message.Chat.ID,
					Text:   fmt.Sprintf("Новый жучара с ID %d создан", id),
				})
			}

			if err != nil {
				log.Println(err.Error())
			}
		}

	}

	log.Println("Shutting down...")
}
