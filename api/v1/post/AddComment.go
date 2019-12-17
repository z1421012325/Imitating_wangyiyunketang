package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context){
	var service user.AddCommentService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.AddComment(c)
	c.JSON(200,res)
}
