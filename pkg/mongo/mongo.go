package mongoRepo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type MongoClient struct {
	client     *mongo.Client
	collection *mongo.Collection
}

type MongoParams struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

func NewMongoParams(host string, username string, password string, port string, database string) *MongoParams {
	return &MongoParams{Host: host, Username: username, Password: password, Port: port, Database: database}
}

func MongoInit() *mongo.Database {
	params := NewMongoParams("localhost", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", params.Username, params.Password, params.Host, params.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("gobot")
}
