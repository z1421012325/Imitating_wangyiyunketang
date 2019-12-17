package middleware

import (
	"demos/model"
	"demos/serialize"
	"fmt"
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

		//user := s.Get("admin")
		//fmt.Println("存在useradmin",user)
		//if user != nil{
		//	fmt.Println("\n user存在,通过...")
		//	if _, ok := user.(model.Admin); ok {
		//		fmt.Println("\n auth通过")
		//		c.Next()
		//		return
		//	}
		//}

		admin := s.Get("admin")
		fmt.Println("admin值为  >> ",admin)

		admin_id := s.Get("admin_id")
		fmt.Println("admin_id值为  >> ",admin_id)

		//if adminid != nil {
		//	_, err := model.GetAdminUser(adminid)
		//	if err == nil {
		//		fmt.Println("\n auth通过")
		//		c.Next()
		//		return
		//	}
		//}


		c.JSON(200, serialize.CheckLogin())
		c.Abort()
	}
}