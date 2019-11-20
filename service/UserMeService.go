package service

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)



func UserMeService(c *gin.Context)*serialize.Response{
	uid := GetUserId(c)
	sql := ""


	DB.DB.Raw(sql,uid).Scan()


	return serialize.Res(nil,"")
}