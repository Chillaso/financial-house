package main

import (
	"github.com/Chillaso/financial-house/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	router.GET("/book/:year/:month", getBookByYearAndMonth)

	router.Run()
}

func getBookByYearAndMonth(c *gin.Context){
	year, err := strconv.Atoi(c.Param("year"))
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		print(err)
	}
	book, err := service.GetBookByYearAndMonth(year, month)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, book)
}
