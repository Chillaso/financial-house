package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterControllers(router *gin.Engine){
	router.GET("/health", health)

	router.GET("/entries", FindAll)
	router.GET("/entry/:year/:month", FindEntriesByYearAndMonth)
	router.POST("/entry", Insert)
	router.PUT("/entry/:id", Update)
	router.DELETE("/entry/:id", Remove)

	router.GET("/entry/result/:year/:month", GetResultByYearAndMonth)
}

func health(context *gin.Context){
	context.JSON(http.StatusOK, "UP")
}