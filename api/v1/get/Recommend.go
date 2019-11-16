package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Recommend(c *gin.Context){
	res := service.RecommendService(c)
	c.JSON(200,res)
}
