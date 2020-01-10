package post

import (
	"demos/serialize"
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func AdoptVideo(c *gin.Context)  {

	var service admin.AdoptVideoService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.AdoptVideo(c)
	c.JSON(200,res)
}
