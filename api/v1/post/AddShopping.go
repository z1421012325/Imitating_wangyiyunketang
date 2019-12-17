package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func AddShopping(c *gin.Context){
	var service user.AddShoppingService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.AddShopping(c)
	c.JSON(200,res)
}
