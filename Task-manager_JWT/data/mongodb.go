package data

import (
	"context"
	"log"
	"time"

	"github.com/Bisrat/task-manager/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func InitDB(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return &errors.ConnectionError{Err: err}
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return &errors.ConnectionError{Err: err}
	}

	db = client.Database("taskmanager")
	log.Println("Connected to MongoDB!")
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

func CloseDB() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}
