package command

import (
	"das-frama/zhukbot-tg/pkg/bot"
	"das-frama/zhukbot-tg/pkg/config"
	"das-frama/zhukbot-tg/pkg/txtdb"
	"fmt"
	"strconv"
)

// Result is a result of command process.
type Result struct {
	Text      string
	PhotoURL  string
	PhotoData []byte
}

var commandMap = map[string]func(*bot.Message, config.Config) (Result, error){
	"start": start,
	"ping":  ping,
}

// Process handles the command and returns a response struct.
func Process(cmd string, m *bot.Message, cfg config.Config) (Result, error) {
	fn, ok := commandMap[cmd]
	if !ok {
		return Result{}, fmt.Errorf("wrong command")
	}

	return fn(m, cfg)
}

func start(message *bot.Message, cfg config.Config) (Result, error) {
	var result Result

	text := ""
	// Insert in db.
	switch message.Chat.Type {
	case "private":
		if message.From.ID == cfg.AdminUserID {
			text = "Привет админ. Бот запущен."
		} else {
			text = "Привет! К сожалению, бот не умеет работать в одиночных чатах."
		}
	case "supergroup":
		ok := cfg.TxtDB.HasRecord("chats.txt", "id", strconv.FormatInt(int64(message.Chat.ID), 10))
		if ok {
			text = "Жукобот уже запущен внутри группы."
		} else {
			err := cfg.TxtDB.Insert("chats.txt", txtdb.Chat{
				ID:            message.Chat.ID,
				Type:          message.Chat.Type,
				Title:         message.Chat.Title,
				Username:      message.Chat.Username,
				FirstName:     message.Chat.FirstName,
				LastName:      message.Chat.LastName,
				SlowModeDelay: message.Chat.SlowModeDelay,
			})
			if err == nil {
				text = "Жукобот запущен внутри группы."
			} else {
				text = err.Error()
			}
		}
	}

	result.Text = text

	return result, nil
}

func ping(message *bot.Message, cfg config.Config) (Result, error) {
	return Result{
		Text: "хуинг",
	}, nil
}
