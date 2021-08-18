package command

import (
	"das-frama/zhukbot-tg/pkg/bot"
	"das-frama/zhukbot-tg/pkg/config"
	"fmt"
)

// Result is a result of command process.
type Result struct {
	Text      string
	PhotoURL  string
	PhotoData []byte
}

var commandMap = map[string]func(*bot.Message, config.Config) (Result, error){
	"start": start,
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
		text = "Жукобот запущен внутри группы."
	}

	result.Text = text

	return result, nil
}
