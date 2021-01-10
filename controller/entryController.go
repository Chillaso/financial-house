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

func FindAll(context *gin.Context){
	entries, err := service.FindAll()
	if err != nil {
		log.Panicln(err)
	} else{
		context.JSON(http.StatusOK, entries)
	}
}

func FindEntriesByYearAndMonth(context *gin.Context){
	year, month := getYearAndMonthParams(context)

	entries, err := service.FindEntriesByYearAndMonth(year, month)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments{
			context.JSON(http.StatusNotFound, entries)
		} else{
			context.JSON(http.StatusInternalServerError, entries)
		}
	} else{
		context.JSON(http.StatusOK, entries)
	}
}

func Insert(context *gin.Context) {
	var entry model.Entry
	err :=  context.Bind(&entry)
	if  err == nil {
		if insertedBookID, err := service.Insert(&entry); err == nil && insertedBookID != nil {
			context.JSON(http.StatusOK, "OK")
		} else {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, -1)
		}
	} else {
		log.Println(err)
		context.JSON(http.StatusBadRequest, err)
	}
}

func Update(context *gin.Context) {
	entryId := context.Param("id")
	var entry model.Entry
	err := context.Bind(&entry)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, err)
	}else{
		err = service.Update(entryId, &entry)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, err)
		} else{
			context.JSON(http.StatusOK, "OK")
		}
	}
}

func Remove(context *gin.Context) {
	entryId := context.Param("id")
	err := service.Remove(entryId)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, err)
	} else{
		context.JSON(http.StatusOK, "OK")
	}
}

func GetResultByYearAndMonth(context *gin.Context){
	year, month := getYearAndMonthParams(context)
	debit, asset, err := service.GetResultByYearAndMonth(year, month)
	if err != nil {
		log.Println(err)
	}
	print(debit, asset)
}

func getYearAndMonthParams(context *gin.Context) (int, int){
	year, err := strconv.Atoi(context.Param("year"))
	if err != nil {
		log.Panicln(err)
	}
	month, err := strconv.Atoi(context.Param("month"))
	if err != nil {
		log.Panicln(err)
	}
	return year, month
}