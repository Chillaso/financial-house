package repository

import (
	"context"
	"github.com/Chillaso/financial-house/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const (
	DATABASE          = "financialHouse"
	COLLECTION 		  = "book"
	MONGOFINANCIALURI = "MONGO_FINANCIAL_URI"
)

var db *mongo.Database

func GetBookByYearAndMonth(year int, month int) (model.Book, error) {

	query := bson.D{{"year", year}, {"months.month", month}}

	result := db.Collection(COLLECTION).FindOne(context.TODO(), query)
	book := model.Book{}
	err := result.Decode(&book)
	if err != nil {
		return book, err
	} else{
		return book, nil
	}
}

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv(MONGOFINANCIALURI)))
	if err != nil {
		log.Panicln(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Panicln(err)
	}
	db = client.Database(DATABASE)
}

