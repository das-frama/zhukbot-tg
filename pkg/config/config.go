package config

import (
	"encoding/json"
	"os"
)

// Config main app configuration struct.
type Config struct {
	Bot struct {
		Mode string `json:"mode"`
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

	return config, nil
}
