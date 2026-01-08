package utils

import (
	"errors"
	"os"
)

func GetTelegramTokenFromEnv() (string, error) {
	token := os.Getenv("TELEGRAM_TOKEN_FLOWER_PATIO")
	if token == "" {
		return "", errors.New("there is no TELEGRAM_TOKEN in .env file")
	}

	return token, nil
}
