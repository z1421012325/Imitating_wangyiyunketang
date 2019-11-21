package post

import (
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

func Modify(c *gin.Context){
	var service service.ModifyService
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(200,serialize.ParamErr("",nil))
		return
	}

	res := service.Modify(c)
	c.JSON(200,res)

}
