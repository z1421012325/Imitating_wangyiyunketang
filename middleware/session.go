package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

func Session() gin.HandlerFunc{
	// cookies,memcache,mongodb,redis
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRE")))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("gin-session", store)

}

//	server.GET("set/session", func(c *gin.Context) {
//		s := sessions.Default(c)
//		s.Set("info","123")
//		s.Save()
//		c.JSON(200,nil)
//	})
//
//	server.GET("get/session", func(c *gin.Context) {
//		s := sessions.Default(c)
//		value := s.Get("info")
//		fmt.Println(value)
//		c.JSON(200,value)
//	})