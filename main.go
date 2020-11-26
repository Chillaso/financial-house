package main

import (
	"github.com/Chillaso/financial-house/service"
	"github.com/gin-gonic/gin"
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
	book := service.GetBookByYearAndMonth(year, month)
	c.JSON(http.StatusOK, book)
}
