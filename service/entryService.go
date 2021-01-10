package service

import (
	"github.com/Chillaso/financial-house/model"
	"github.com/Chillaso/financial-house/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ASSETTYPE = 1
	DEBITTYPE = 0
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

func Update(entryId string, entry *model.Entry) error{
	id, err := primitive.ObjectIDFromHex(entryId)
	if err != nil {
		return err
	} else{
		return repository.Update(id, entry)
	}
}

func Remove(entryId string) error {
	id, err := primitive.ObjectIDFromHex(entryId)
	if err != nil{
		return err
	} else{
		return repository.Remove(id)
	}
}

func GetResultByYearAndMonth(year int, month int) (float32, float32, error){
	entries, err := FindEntriesByYearAndMonth(year, month)
	if err != nil {
		return -1, -1, err
	}
	var debit float32
	var asset float32
	for _, entry := range entries {
		if entry.Type == DEBITTYPE {
			debit += entry.Amount
		} else if entry.Type == ASSETTYPE {
			asset += entry.Amount
		}
	}
	return debit, asset, nil
}
