package controller

import (
	"github.com/Chillaso/financial-house/model"
	"github.com/Chillaso/financial-house/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

func GetAllBooks(context *gin.Context){
	books, err := service.GetAllBooks()
	if err != nil {
		log.Panicln(err)
	} else{
		context.JSON(http.StatusOK, books)
	}
}

func GetBookByYearAndMonth(context *gin.Context){
	year, err := strconv.Atoi(context.Param("year"))
	if err != nil {
		log.Panicln(err)
	}
	month, err := strconv.Atoi(context.Param("month"))
	if err != nil {
		log.Panicln(err)
	}

	book, err := service.GetBookByYearAndMonth(year, month)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments{
			context.JSON(http.StatusNotFound, book)
		} else{
			context.JSON(http.StatusInternalServerError, book)
		}
	} else{
		context.JSON(http.StatusOK, book)
	}
}

func AddBook(context *gin.Context) {
	var book model.Book
	err :=  context.Bind(&book)
	if  err == nil {
		if insertedBookID, err1 := service.AddBook(&book); err1 != nil && insertedBookID != nil {
			context.JSON(http.StatusOK, insertedBookID.InsertedID)
		}else {
			log.Println(err1)
			context.JSON(http.StatusInternalServerError, -1)
		}
	} else {
		log.Println(err)
		context.JSON(http.StatusBadRequest, err)
	}
}

func AddEntry(context *gin.Context) {

}

func ModifyEntry(context *gin.Context) {

}

func RemoveEntry(context *gin.Context) {

}