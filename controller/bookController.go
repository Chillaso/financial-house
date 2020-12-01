package controller

import (
	"github.com/Chillaso/financial-house/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

func GetBookByYearAndMonth(c *gin.Context){
	year, err := strconv.Atoi(c.Param("year"))
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		log.Println(err)
	}
	book, err := service.GetBookByYearAndMonth(year, month)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments{
			c.JSON(http.StatusNotFound, book)
		} else{
			c.JSON(http.StatusInternalServerError, book)
		}
	} else{
		c.JSON(http.StatusOK, book)
	}
}
