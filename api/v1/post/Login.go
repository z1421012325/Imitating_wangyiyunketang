package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)


// 登录
func Login(c* gin.Context) {
	var service user.LoginService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.Login(c)
	c.JSON(200,res)
}