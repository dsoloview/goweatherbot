package main

import (
	mongoRepo "github.com/dsoloview/gobot/pkg/mongo"
	"github.com/dsoloview/gobot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	database := mongoRepo.MongoInit()
	bot.Debug = false

	telegramBot := telegram.NewBot(bot, database)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
