package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ModifyInfo(c *gin.Context){
	var service user.ModifyService
	if err := c.ShouldBind(&service);err!=nil{
		c.JSON(200,serialize.ParamErr("",nil))
		return
	}

	res := service.Modify(c)
	c.JSON(200,res)

}
