package service

import (
	"github.com/Chillaso/financial-house/model"
	"github.com/Chillaso/financial-house/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllBooks() ([]model.Book, error){
	return repository.GetAllBooks()
}

func GetBookByYearAndMonth(year int, month int) (model.Book, error) {
	return repository.GetBookByYearAndMonth(year, month)
}

func AddBook(book *model.Book) (*mongo.InsertOneResult, error){
	return repository.AddBook(book)
}

