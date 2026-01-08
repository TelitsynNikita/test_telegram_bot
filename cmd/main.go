package main

import (
	"log"

	"github.com/TelitsynNikita/test_telegram_bot/internal/handler"
	"github.com/TelitsynNikita/test_telegram_bot/internal/repository"
	"github.com/TelitsynNikita/test_telegram_bot/internal/service"
	"github.com/TelitsynNikita/test_telegram_bot/internal/utils"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token, err := utils.GetTelegramTokenFromEnv()
	if err != nil {
		log.Fatal("Error getting telegram token")
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(token, services)

	handlers.InitHandlers()

	log.Fatal(handlers.Server.Start())
}
