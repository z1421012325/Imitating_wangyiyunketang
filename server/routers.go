package server

import (
	_ "demos/DB"

	"github.com/gin-gonic/gin"

	v1get "demos/api/v1/get"
	v1post "demos/api/v1/post"
	"demos/middleware"
)


func NewRouter() *gin.Engine{
	Router := gin.Default()

	// 心跳检测
	Router.GET("ping",func (c *gin.Context){
		c.JSON(200,gin.H{
			"msg":"ping",
		})
	})

	// 中间件 跨域最前,session,auth
	Router.Use(middleware.Cors())
	Router.Use(middleware.Session())
	Router.Use(middleware.Auth())


	
	// 版本迭代
	v1 := Router.Group("/api/v1")
	{
		// 用户注册
		v1.POST("registry/user",v1post.RegistryUser)
		// 用户登录
		v1.POST("login",v1post.Login)
		// 课程页面
		v1.GET("course/introduction/:cid",v1get.Introduction)
		// 课程目录
		v1.GET("course/coursedetail/:cid",v1get.Coursedetail)

		// 中间件,保护登录
		v1.Use(middleware.AuthLogin())
		{
			// 退出
			v1.POST("logout",v1post.Logout)
		}

	}



























	return Router

}