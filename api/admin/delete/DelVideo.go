package delete

import (
	"demos/serialize"
	"demos/service/admin"
	"github.com/gin-gonic/gin"
)

func DelVideo(c *gin.Context){
	var service admin.DelVideoServic
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}

	res := service.DelVideo(c)
	c.JSON(200,res)

}
