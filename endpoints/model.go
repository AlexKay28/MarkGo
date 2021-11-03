package endpoints

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Array1 []int  `json:"array1"`
	Array2 []int  `json:"array2"`
}

func CalculateModel(c *gin.Context) {
  var data Data
  if err := c.BindJSON(&data); err != nil {
    c.AbortWithError(400, err)
    return
  }

  var result []int
  for i := 0; i < len(data.Array1); i++ {
    result = append(result, data.Array1[i] * data.Array2[i])
  }
  c.JSON(http.StatusOK, result)
}
