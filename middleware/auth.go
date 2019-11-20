package middleware

import (
	"demos/model"
	"demos/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)



func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		s := sessions.Default(c)

		uid := s.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				s.Set("user", user)
				_ = s.Save()
			}
		}
		c.Next()
	}
}


// 登录用户权限控制
func AuthLogin() gin.HandlerFunc{
	return func(c *gin.Context) {

		s := sessions.Default(c)
		user := s.Get("user")
		if user != nil{
			if _, ok := user.(model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serialize.CheckLogin())
		c.Abort()
	}
}





// 管理员权限控制
func AuthAdminLogin() gin.HandlerFunc{

	return func(c *gin.Context) {
		s := sessions.Default(c)

		user := s.Get("admin")
		if user != nil{
			if _, ok := user.(model.Admin); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serialize.CheckLogin())
		c.Abort()
	}
}