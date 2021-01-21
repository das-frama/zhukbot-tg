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
			if update.Message.Command() == "start" {
				err := db.CreateUser(postgres.User{
					ID:                      update.Message.From.ID,
					FirstName:               update.Message.From.FirstName,
					LastName:                update.Message.From.LastName,
					Username:                update.Message.From.Username,
					LanguageCode:            update.Message.From.LanguageCode,
					CanJoinGroups:           update.Message.From.CanJoinGroups,
					CanReadAllGroupMessages: update.Message.From.CanReadAllGroupMessages,
				})
				if err != nil {
					log.Println(err.Error())
				}
			}
		}

	}

	log.Println("Shutting down...")
}
