package server

import (
	_ "demos/DB"
	_ "demos/conf"

	"demos/middleware"

	"github.com/gin-gonic/gin"
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
		v1.GET("")

	}



























	return server

}