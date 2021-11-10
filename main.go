package main

import (
	"github.com/gin-gonic/gin"
	"github.com/AlexKay28/MarkGo/endpoints"
	"github.com/AlexKay28/MarkGo/utils"
)

func main() {
	router := gin.Default()
	router.GET("/home", endpoints.PrintMessage)
	router.GET("/transmat", utils.OptrainEventsFromCsv)
	router.Run("localhost:5566")
}
