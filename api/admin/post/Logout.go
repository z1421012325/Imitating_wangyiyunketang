package post

import (
	"demos/model"
	"demos/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context)  {

	s := sessions.Default(c)
	admin := s.Get("admin")
	if admin != nil{
		if _, ok := admin.(*model.Admin); ok {
			s.Clear()
			_ = s.Save()
			c.JSON(200,serialize.Res(nil,"登出成功"))
		}
	}

	c.JSON(404,nil)

}
