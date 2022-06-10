package main

import (
	"github.com/joho/godotenv"
	"gobot/pkg/gismeteo"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	gismeteo.SearchCity("Москва")

	//bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	//if err != nil {
	//	log.Panic(err)
	//}
	//bot.Debug = false
	//
	//telegramBot := telegram.NewBot(bot)
	//if err := telegramBot.Start(); err != nil {
	//	log.Fatal(err)
	//}
}
