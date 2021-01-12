package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config main app configuration struct.
type Config struct {
	Bot struct {
		Token string
		Mode  string `json:"mode"`
	}
	DB struct {
		URL string
	}
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
	config.Bot.Token = os.Getenv("BOT_TOKEN")
	if config.Bot.Token == "" {
		return config, fmt.Errorf("$BOT_TOKEN must be set")
	}

	// Database URL
	config.DB.URL = os.Getenv("DATABASE_URL")
	if config.DB.URL == "" {
		return config, fmt.Errorf("$DATABASE_URL must be set")
	}

	return config, nil
}
