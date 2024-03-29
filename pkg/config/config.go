package config

import (
	"das-frama/zhukbot-tg/pkg/txtdb"
	"encoding/json"
	"fmt"
	"os"
)

// Config main app configuration struct.
type Config struct {
	AdminUserID int

	Bot struct {
		Token string
		Mode  string `json:"mode"`
	}

	DB struct {
		Path string `json:"path"`
	}
	TxtDB *txtdb.DB
}

// LoadConfig creates a struct from file.
func LoadConfig(filepath string) (Config, error) {
	var config Config
	// Open file.
	file, err := os.Open(filepath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Parse config.
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}

	// Get env.
	// Get bot api token from env.
	env, err := GetEnv()
	if err != nil {
		return config, fmt.Errorf(".env file does not exists, make sure you copy .env.example to .env")
	}
	config.Bot.Token = env.BotToken
	config.AdminUserID = env.AdminUserID

	return config, nil
}
