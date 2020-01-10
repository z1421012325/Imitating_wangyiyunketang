package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func GetOss(c *gin.Context)  {
	var service user.GetOssService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.GetOss()
	c.JSON(200,res)
}
