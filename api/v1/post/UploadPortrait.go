package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func UploadPortrait(c *gin.Context)  {
	var service user.UploadPortraitService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.UploadPortrait(c)
	c.JSON(200,res)

}
