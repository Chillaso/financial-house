package service

import (
	"github.com/Chillaso/financial-house/model"
	"github.com/Chillaso/financial-house/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll() ([]model.Entry, error){
	return repository.FindAll()
}

func FindEntriesByYearAndMonth(year int, month int) ([]model.Entry, error) {
	return repository.FindEntriesByYearAndMonth(year, month)
}

func Insert(entry *model.Entry) (*mongo.InsertOneResult, error){
	return repository.Insert(entry)
}

func Update(entry *model.Entry) error{
	return repository.Update(entry)
}

func Remove(entryId string) error {
	id, err := primitive.ObjectIDFromHex(entryId)
	if err != nil{
		return err
	} else{
		return repository.Remove(id)
	}
}

