package post

import (
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

func Login(c* gin.Context) {
	var service service2.LoginService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
	}

	res := service.Login(c)
	c.JSON(200,res)


}