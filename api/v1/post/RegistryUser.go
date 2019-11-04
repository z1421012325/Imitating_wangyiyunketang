package post

import (
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)


// 注册(身份注册,学生,老师)
func RegistryUser(c *gin.Context){
	var service service.RegistryUserService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
	}

	res := service.Register()
	c.JSON(200,res)

}