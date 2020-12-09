package repository

import (
	"context"
	"github.com/Chillaso/financial-house/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const (
	DATABASE          = "financialHouse"
	COLLECTION 		  = "entry"
	MONGOFINANCIALURI = "MONGO_FINANCIAL_URI"
)

var (
	db *mongo.Database
)

func FindAll() ([]model.Entry, error){
	entries := make([]model.Entry, 0, 2)
	//TODO: Should have some method to page results
	cursor, err := db.Collection(COLLECTION).Find(context.Background(), bson.D{}, options.Find().SetLimit(int64(10)))
	if cursor != nil && err == nil {
		ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
		for cursor.Next(ctx) {
			entry := model.Entry{}
			err = cursor.Decode(&entry)
			entries = append(entries, entry)
		}
	}
	return entries, err
}

func FindEntriesByYearAndMonth(year int, month int) ([]model.Entry, error) {
	var entries []model.Entry
	query := bson.D{{"year", year}, {"month",month}}
	cursor, err := db.Collection(COLLECTION).Find(context.Background(), query, options.Find().SetLimit(int64(100)))
	if cursor != nil && err == nil {
		ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
		for cursor.Next(ctx) {
			entry := model.Entry{}
			err = cursor.Decode(&entry)
			entries = append(entries, entry)
		}
	}
	return entries, err
}

func Insert(entry *model.Entry) (*mongo.InsertOneResult, error){
	return db.Collection(COLLECTION).InsertOne(context.Background(), entry)
}

func Update(id primitive.ObjectID, entry *model.Entry) error {
	_, err := db.Collection(COLLECTION).ReplaceOne(context.Background(), bson.D{{"_id", id}}, entry)
	return err
}

func Remove(entryId primitive.ObjectID) error {
	_, err := db.Collection(COLLECTION).DeleteOne(context.Background(), bson.D{{"_id", entryId}})
	return err
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

