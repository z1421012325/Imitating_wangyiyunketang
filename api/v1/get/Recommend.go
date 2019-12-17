package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func Recommend(c *gin.Context){
	res := user.RecommendService(c)
	c.JSON(200,res)
}
