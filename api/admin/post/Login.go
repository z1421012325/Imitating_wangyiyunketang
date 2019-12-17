package post

import (
	"demos/serialize"
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var service admin.LoginService
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.Login(c)
	c.JSON(200,res)
}