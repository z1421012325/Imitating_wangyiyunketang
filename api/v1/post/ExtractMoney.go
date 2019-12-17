package post

import (
	"demos/serialize"
	"demos/service/user"
	"github.com/gin-gonic/gin"
)

func ExtractMoney(c *gin.Context){
	var service user.ExtractMoneyService
	if err := c.ShouldBind(&service);err != nil{
		c.JSON(200,serialize.ParamErr("",err))
		return
	}
	res := service.ExtractMoney(c)
	c.JSON(200,res)
}
