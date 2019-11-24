package server

import (
	_ "demos/DB"

	v1get "demos/api/v1/get"
	v1post "demos/api/v1/post"
	v1del "demos/api/v1/delete"
	"demos/middleware"

	"github.com/gin-gonic/gin"
)


func NewRouter() *gin.Engine{
	Router := gin.Default()

	// 心跳检测
	Router.GET("ping",v1get.Ping)

	// 中间件 跨域最前,session,auth
	Router.Use(middleware.Cors())
	Router.Use(middleware.Session())
	Router.Use(middleware.Auth())


	// 版本迭代
	v1 := Router.Group("/api/v1")
	{

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
		// 开始学习,要对是否登录用户检测,或者课程价格为0or不为0检测(是否购买检测)
		v1.GET("course/standard/:cid",v1get.Standard)
		// 课程推荐(老师本人还是根据标签推荐有前端决定,后端输出 degree)
		v1.GET("recommend/:cid",v1get.Recommend)


		// search 搜索
		v1.GET("search",v1get.Search)
		// 查询tag
		v1.GET("search/tag",v1get.SearchTag)


		// 所有人都能查看到的学生信息
		v1.GET("student/:uid",v1get.Student)
		// 所有人都能看到的老师信息(老师介绍页面,能看到所教学的课程,无需登录即可看到)
		v1.GET("instructorinfo/:uid",v1get.InstructorInfo)
		// 该老师正在教学课程
		v1.GET("curriculum/now/:uid",v1get.NowShowlist)


		// 用户注册
		v1.POST("registry/user",v1post.RegistryUser)
		// 用户登录
		v1.POST("login",v1post.Login)


		{// 中间件,保护登录
			v1.Use(middleware.AuthLogin())
			// 查看学习的视频
			v1.GET("show/study",v1get.ShowStudy)
			// 查看个人信息
			v1.GET("user/me",v1get.UserMe)
			// 修改个人信息 和查看个人信息配合
			v1.POST("user/modify/info",v1post.ModifyInfo)
			// 修改密码
			v1.POST("user/modify/pswd",v1post.ModifyPswd)

			// 添加视频收藏
			v1.POST("add/collection",v1post.AddCollection)
			// 查看收藏视频
			v1.GET("show/collection",v1get.ShowCollection)
			// 取消收藏

			// 上传视频
			// 查看视频信息
			// 修改视频信息

			// 下架视频
			v1.DELETE("del/curriculum",v1del.DelCurriculum)
			// 查看下架的视频
			// 回复下架视频

			// 查看课程目录中的视频
			// 增加课程目录中的视频

			// 发表评论
			// 查看评论
			// 删除评论

			// 添加购物车,不是直接订单
			// 购物车下单状态更改 添加订单

			// 查看课程购买记录(老师)
			// 查看拥有金额(老师)
			// 提成金额(老师)
			// 提成记录(老师)


			// 上传用户头像


			// 退出
			v1.POST("logout",v1post.Logout)
		}



		// 后台管理,或者使用另一个web服务开管理,不过记得模型要一样
		{
			v1.Use(middleware.AuthAdminLogin())
			// 查看视频和总个数
			// 根据天数查看当天视频上传个数 七天为期
			// 根据天数查看当天注册的人数(学生or老师or全部)
			// 查看总注册人数(老师or学生or全部)

			// 删除视频
			// 删除评论

			// 根据天数查看当天下单金额
			// 根据月份查看当月下单金额
			// 总共下单金额

			// 站点提成金额总数
			// 根据天数查看当天提成金额
			// ...

			// 被老师提取出去的金额总数
			// 根据天数查看被老师提取出去的金额总数
			// ...

		}

	}



	return Router

}