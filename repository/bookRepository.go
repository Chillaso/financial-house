package repository

import (
	"context"
	"github.com/Chillaso/financial-house/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetBookByYearAndMonth(year int, month int) (model.Book, error) {
	client := getClient()
	ctx, _ :=context.WithTimeout(context.Background(), 10 * time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return model.Book{}, nil
}

func getClient() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
