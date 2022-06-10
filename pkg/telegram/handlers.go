package telegram

import (
	"fmt"
	"github.com/dsoloview/gobot/pkg/openweather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) tgbotapi.MessageConfig {
	text := ""
	weather, err := openweather.GetWeather(message.Text)
	if err != nil {
		text = err.Error()
	} else {
		text = fmt.Sprintf("Температура в городе %v: %v градусов Цельсия. \n По ощущениям: %v. \n %s", weather.City.City, weather.Temperature, weather.Temperature, weather.Description)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID

	return msg
}

func (b *Bot) handleCommand(message *tgbotapi.Message) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	switch message.Command() {
	case "start":
		msg.Text = "Привет! Это бот, который умеет отправлять текущую погоду по заданному городу☀️ \n Например, введи \"Москва\" и получишь температуру в Москве"
	}

	return msg
}
