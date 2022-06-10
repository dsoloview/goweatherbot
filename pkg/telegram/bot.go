package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {

	log.Printf("Bot started on username @%s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	msg := tgbotapi.MessageConfig{}

	for update := range updates {

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			msg = b.handleCommand(update.Message)
		} else {
			msg = b.handleMessage(update.Message)
		}
		err := b.sendResponse(msg)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func (b *Bot) sendResponse(msg tgbotapi.MessageConfig) error {
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
