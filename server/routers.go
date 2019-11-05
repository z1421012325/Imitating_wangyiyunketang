package server

import (
	_ "demos/conf"
	_ "demos/DB"

	"github.com/gin-gonic/gin"

	"demos/middleware"
	v1post "demos/api/v1/post"

)


func NewRouter() *gin.Engine{
	server := gin.Default()

	// 心跳检测
	server.GET("ping",func (c *gin.Context){
		c.JSON(200,gin.H{
			"msg":"ping",
		})
	})


	// 中间件 跨域最前,session,auth
	server.Use(middleware.Cors())
	server.Use(middleware.Session())
	server.Use(middleware.Auth())



	
	
	// 版本迭代
	v1 := server.Group("/api/v1")
	{
		// 用户注册
		v1.POST("registry/user",v1post.RegistryUser)

		// 用户登录
	}



























	return server

}