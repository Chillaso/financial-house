package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterControllers(router *gin.Engine){
	router.GET("/health", health)

	router.GET("/book", GetAllBooks)
	router.GET("/book/:year/:month", GetBookByYearAndMonth)
	router.POST("/book", AddBook)

	router.POST("/book/:year/:month", AddEntry)
	router.PUT("/book/:year/:month/:entryId", ModifyEntry)
	router.DELETE("/book/:year/:month/:entryId", RemoveEntry)

}

func health(context *gin.Context){
	context.JSON(http.StatusOK, "UP")
}