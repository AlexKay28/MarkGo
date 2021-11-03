package main

import (
	"github.com/gin-gonic/gin"
	"github.com/AlexKay28/golang_web_ml/endpoints"
	"github.com/AlexKay28/golang_web_ml/models"
)

func main() {
	models.BuildModel()

	router := gin.Default()
	router.GET("/model", endpoints.CalculateModel)
	router.GET("/home", endpoints.PrintMessage)
	router.Run("localhost:5566")
}
