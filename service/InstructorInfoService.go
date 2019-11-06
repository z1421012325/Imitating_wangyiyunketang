package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

func InstructorInfoService(c *gin.Context) *serialize.Response {
	uid := c.Param("uid")
	var user model.User
	DB.DB.Where("u_id = ?",uid).First(&user)

	return serialize.Res(user,"")
}