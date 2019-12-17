package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

func HaveMoneyService(c *gin.Context)*serialize.Response{
	uid := service.GetUserId(c)

	var data model.Money
	DB.DB.Where("u_id = ?",uid).First(&data)

	return serialize.Res(data,"")
}
