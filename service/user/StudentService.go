package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

func StudentService(c *gin.Context) *serialize.Response {
	uid := c.Param("uid")

	var user model.User
	DB.DB.Select("u_id,nickename,r_id,portrait,create_at").
		Where("u_id = ? and r_id = ?",uid,model.Student).
		First(&user)

	user.CompletionToOssUrl()

	return serialize.Res(user,"")
}
