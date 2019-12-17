package get

import (
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func Comment(c *gin.Context){
	res := user.CommentService(c)
	c.JSON(200,res)
}