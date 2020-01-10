package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func SaveNewVideo(c *gin.Context){
	var service user.SavaNewVideoService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.SavaNewVideo(c)
	c.JSON(200,res)
}
