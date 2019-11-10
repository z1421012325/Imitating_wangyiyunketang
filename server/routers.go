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
		v1.GET("course/coursedetail/:cid",v1get.CourseDetail)
		// 课程评论
		v1.GET("course/comment/:cid",v1get.Comment)
		// 展示所含有的老师
		v1.GET("instructor/all",v1get.AllInstructorInfo)
		// 课程目录
		v1.GET("catalog/:cid",v1get.Catalog)



		// 所有人都能看到的老师信息(老师介绍页面,能看到所教学的课程,无需登录即可看到)
		v1.GET("instructorinfo/:uid",v1get.InstructorInfo)
		// 正在教学课程
		v1.GET("curriculum/now/:uid",v1get.NowShowlist)



		// 中间件,保护登录
		v1.Use(middleware.AuthLogin())
		{
			// 退出
			v1.POST("logout",v1post.Logout)
		}

	}



























	return Router

}