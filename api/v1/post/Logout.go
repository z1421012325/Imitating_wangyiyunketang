package post

import (
	"demos/model"
	"demos/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context){
	s := sessions.Default(c)
	user := s.Get("user")
	if user != nil{
		if _, ok := user.(*model.User); ok {
			s.Clear()
			_ = s.Save()
			//c.JSON(200, serialize.Response{
			//	Code: 0,
			//	Msg:  "登出成功",
			//})
			c.JSON(200,serialize.Res(nil,"登出成功"))
		}
	}

	c.JSON(404,nil)
}