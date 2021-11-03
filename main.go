package main

import (
	"github.com/gin-gonic/gin"
	"github.com/AlexKay28/MarkGo/endpoints"
	"github.com/AlexKay28/MarkGo/models"
)

func main() {
	models.BuildModel()

	router := gin.Default()
	router.GET("/model", endpoints.CalculateModel)
	router.GET("/home", endpoints.PrintMessage)
	router.Run("localhost:5566")
}
