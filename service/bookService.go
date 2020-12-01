package service

import (
	"github.com/Chillaso/financial-house/model"
	"github.com/Chillaso/financial-house/repository"
)

func GetBookByYearAndMonth(year int, month int) (model.Book, error) {
	book, err := repository.GetBookByYearAndMonth(year, month)
	return book, err
}
