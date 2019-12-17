package post

import (
	"demos/serialize"
	"github.com/gin-gonic/gin"
	"demos/service/admin"
)


func AdoptUser(c *gin.Context){
	var service admin.AdoptUserServic
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.AdoptUser()
	c.JSON(200,res)
}
