package endpoints

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func PrintMessage(c *gin.Context) {
    message := "This is message from endpoint HOME"
    c.JSON(http.StatusOK, message)
}
