package get

import (
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Comment(c *gin.Context){
	res := service.CommentService(c)
	c.JSON(200,res)
}