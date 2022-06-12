package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TelegramUser struct {
	ID         primitive.ObjectID `bson:"_id"`
	TelegramId int64              `bson:"telegram_id"`
	Username   string             `bson:"username"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	Messages   []Message          `bson:"messages"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id"`
	Message   string             `bson:"message"`
	MessageId int                `bson:"message_id"`
	CreatedAt time.Time          `bson:"created_at"`
}

type db struct {
	collection *mongo.Collection
}

func NewDb(database *mongo.Database, collection string) *db {
	return &db{collection: database.Collection(collection)}
}

func (d *db) Create(ctx context.Context, user TelegramUser) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	return "", err
}

func (d *db) FindByTelegramId(telegramId int64) bool {
	filter := bson.M{"telegram_id": telegramId}
	result := d.collection.FindOne(context.TODO(), filter)

	if result.Err() != nil {
		return false
	}
	return true
}

func (d *db) SaveMessage(userId int64, message Message) error {
	fmt.Println(message)
	filter := bson.M{"telegram_id": userId}
	change := bson.M{"$push": bson.M{"messages": message}}
	_, err := d.collection.UpdateOne(context.TODO(), filter, change)
	if err != nil {
		return err
	}
	return nil
}
