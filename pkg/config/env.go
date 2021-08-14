package config

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Env struct {
	BotToken     string `env:"BOT_TOKEN"`
	DBConnection string `env:"DB_CONNECTION"`
}

func GetEnv() (Env, error) {
	// Read the entire file.
	file, err := os.Open(".env")
	if err != nil {
		return Env{}, err
	}

	defer file.Close()

	// Prepare struct for populating.
	env := Env{}
	t := reflect.TypeOf(env)
	v := reflect.ValueOf(&env).Elem()

	// Scan .env file.
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		envLine := strings.TrimSpace(scanner.Text())
		// Check if string has '=' rune.
		if !strings.ContainsRune(envLine, '=') || strings.HasPrefix(envLine, "#") {
			continue
		}

		// Split string.
		s := strings.Split(envLine, "=")
		envKey, envValueStr := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])

		// Find a propert struct type to pupulate.
		for i := 0; i < t.NumField(); i++ {
			tag := t.Field(i).Tag.Get("env")
			if tag == envKey {
				// If env value is numeric...
				if envValueInt, err := strconv.ParseInt(envValueStr, 10, 64); err == nil {
					v.Field(i).SetInt(envValueInt)
				} else if envValueBool, err := strconv.ParseBool(envValueStr); err == nil {
					v.Field(i).SetBool(envValueBool)
				} else {
					v.Field(i).SetString(envValueStr)
				}
			}
		}
	}

	return env, nil
}
