package main

import (
	"api/api/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/quotes/:ID", controller.GetById)
	router.GET("/quotes/random/:limit", controller.GetRandom)
	router.POST("/quotes", controller.New)
	router.PUT("/quotes/:ID", controller.Update)
	router.DELETE("/quotes/:ID", controller.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
