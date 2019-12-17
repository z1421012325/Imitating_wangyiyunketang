package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ModifyShoppingStatus(c *gin.Context){
	var service user.ModifyShoppingStatusService
	if err := c.ShouldBind(&service);err!= nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.ModifyShoppingStatus(c)
	c.JSON(200,res)
}