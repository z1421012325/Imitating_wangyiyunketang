package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func SaveAddVideo(c *gin.Context)  {
	var service user.SavaAddVideoService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.SavaAddVideo(c)
	c.JSON(200,res)
}
