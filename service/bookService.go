package service

import (
	"github.com/Chillaso/financial-house/model"
	"github.com/Chillaso/financial-house/repository"
)

func GetAllBooks() ([]model.Book, error){
	return repository.GetAllBooks()
}

func GetBookByYearAndMonth(year int, month int) (model.Book, error) {
	return repository.GetBookByYearAndMonth(year, month)
}

