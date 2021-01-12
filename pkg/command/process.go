package command

import (
	"das-frama/zhukbot-tg/pkg/bot"
	"das-frama/zhukbot-tg/pkg/config"
	"fmt"
)

// Config struct
type commandConfig struct {
	Message   *bot.Message
	AppConfig config.Config
}

// Result is a result of command process.
type Result struct {
	Text      string
	PhotoURL  string
	PhotoData []byte
}

var commandMap = map[string]func(commandConfig) (Result, error){}

// Process handles the command and returns a response struct.
func Process(cmd string, m *bot.Message, cfg config.Config) (Result, error) {
	fn, ok := commandMap[cmd]
	if !ok {
		return Result{}, fmt.Errorf("wrong command")
	}

	return fn(commandConfig{
		Message:   m,
		AppConfig: cfg,
	})
}
