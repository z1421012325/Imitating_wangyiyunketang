package get

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context){
	c.JSON(200,gin.H{
		"msg":"ping",
	})
}
