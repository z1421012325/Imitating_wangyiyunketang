package post

import (
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context){
	var service service2.AddCommentService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.AddComment(c)
	c.JSON(200,res)
}
