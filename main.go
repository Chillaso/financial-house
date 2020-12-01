package main

import (
	"github.com/Chillaso/financial-house/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	controller.RegisterControllers(router)
	router.Run()
	log.Print("FINANCIAL HOUSE SERVICE RUNNING...")
}

