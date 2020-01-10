package post

import (
	"demos/serialize"
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func AdoptComment(c *gin.Context)  {
	var service admin.AdoptCommentService
	if err := c.ShouldBind(&service); err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.AdoptComment(c)
	c.JSON(200,res)
}
