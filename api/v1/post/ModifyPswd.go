package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ModifyPswd(c *gin.Context){
	var service user.ModifyPswdService
	if err := c.ShouldBind(&service); err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.ModifyPswd(c)
	c.JSON(200,res)

}
