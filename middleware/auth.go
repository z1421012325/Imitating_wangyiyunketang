package middleware

import (
	"demos/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func Auth() gin.HandlerFunc{
	
	return func(c *gin.Context) {
		s := sessions.Default(c)

		uid := s.Get("user_id")

		if uid != nil {
			control,ok := s.Get("control").(int)
			if ok{
				if (control >= 2 && control < 10 ){
					control++
					s.Set("control",control)
					_ = s.Save()
					c.Next()
					return
				} else if control >= 10 {
					s.Set("query_control",1)
					_ = s.Save()
				} else {
					control++
					s.Set("query_control",control)
					_ = s.Save()
				}

			}
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
				_ = s.Save()
			}
		}

		c.Next()
	}
}

