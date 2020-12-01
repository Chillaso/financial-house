package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterControllers(router *gin.Engine){
	router.GET("/health", health)
	router.GET("/book/:year/:month", GetBookByYearAndMonth)
}

func health(context *gin.Context){
	context.JSON(http.StatusOK, "UP")
}