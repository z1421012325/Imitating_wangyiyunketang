package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"

	"demos/service"
)



func UserMeService(c *gin.Context)*serialize.Response{
	uid := service.GetUserId(c)
	sql := "select " +
				"u.u_id,u.nickename,u.username,u.status,u.r_id,u.portrait,u.create_at " +
			"from " +
				"users as u " +
			"where " +
				"u.u_id = ?"

	var user model.User
	DB.DB.Raw(sql,uid).Scan(&user)

	user.CompletionToOssUrl()

	return serialize.Res(user,"")
}