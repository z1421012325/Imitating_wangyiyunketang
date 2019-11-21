package post

import (
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

func ModifyPswd(c *gin.Context){
	var service service2.ModifyPswdService
	if err := c.ShouldBind(&service); err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.ModifyPswd(c)
	c.JSON(200,res)

}
