package telegram

import (
	"context"
	"fmt"
	"github.com/dsoloview/gobot/pkg/mongo/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	database *mongo.Database
}

func NewBot(bot *tgbotapi.BotAPI, database *mongo.Database) *Bot {
	return &Bot{bot: bot, database: database}
}

func (b *Bot) Start() error {

	log.Printf("Bot started on username @%s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	msg := tgbotapi.MessageConfig{}
	db := repository.NewDb(b.database, "telegramUsers")

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

		// Save user to database
		if !db.FindByTelegramId(update.Message.From.ID) {
			user := createTelegramUser(update.Message.From)

			_, err = db.Create(context.TODO(), user)
			if err != nil {
				log.Fatal(err)
			}
		}

		// Message save
		message := createMessage(update.Message)
		fmt.Println(message)
		err = db.SaveMessage(update.Message.From.ID, message)
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

func createTelegramUser(chat *tgbotapi.User) repository.TelegramUser {
	return repository.TelegramUser{
		ID:         primitive.NewObjectID(),
		TelegramId: chat.ID,
		Username:   chat.UserName,
		FirstName:  chat.FirstName,
		LastName:   chat.LastName,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Messages:   []repository.Message{},
	}
}

func createMessage(message *tgbotapi.Message) repository.Message {
	return repository.Message{
		ID:        primitive.NewObjectID(),
		Message:   message.Text,
		MessageId: message.MessageID,
		CreatedAt: time.Now(),
	}
}
