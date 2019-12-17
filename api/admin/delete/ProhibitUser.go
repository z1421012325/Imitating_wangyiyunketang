package delete

import (
	"demos/serialize"
	"github.com/gin-gonic/gin"
	"demos/service/admin"
	)
func ProhibitUser(c *gin.Context){
	var service admin.ProhibitUserService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.ProhibitUser(c)
	c.JSON(200,res)



}
