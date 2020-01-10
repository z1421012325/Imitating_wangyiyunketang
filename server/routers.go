package server

import (
	_ "demos/DB"
	_ "demos/util"

	admindel "demos/api/admin/delete"
	adminget "demos/api/admin/get"
	adminpost "demos/api/admin/post"
	v1del "demos/api/v1/delete"
	v1get "demos/api/v1/get"
	v1post "demos/api/v1/post"

	"demos/middleware"

	"github.com/gin-gonic/gin"
)


func NewRouter() *gin.Engine{
	Router := gin.Default()

	Router.GET("ping",v1get.Ping)

	// 中间件 跨域最前,session,auth
	Router.Use(middleware.Cors())
	Router.Use(middleware.Session())
	Router.Use(middleware.Auth())

	// 支付宝回调信息
	Router.POST("zfb/callback",v1post.ZfbCallback)


	// 版本迭代
	v1 := Router.Group("/api/v1")
	{
		// 课程页面
		v1.GET("course/introduction/:cid",v1get.Introduction)
		// 课程目录
		// v1.GET("course/coursedetail/:cid",v1get.CourseDetail)
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
		// 查看点击排行
		v1.GET("show/list/click",v1get.ShowListClick)


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
			// 取消(删除)收藏
			v1.DELETE("del/collection",v1del.DelCollection)


			// 给予aliyun-oss凭证 让前端去aliyun-oss上传
			v1.POST("get/oss/token",v1post.GetOss)
			// 新增视频 保存上传视频(老师)url
			v1.POST("save/new/video",v1post.SaveNewVideo)
			// 后续增加课程目录中的视频
			v1.POST("save/add/video",v1post.SaveAddVideo)
			// 保存用户头像
			v1.POST("save/portrait",v1post.UploadPortrait)


			// 查看视频信息(老师)
			v1.GET("see/teacher/curriculum/video/:cid",v1get.TeacherCurriculumVideo)
			// 修改视频目录信息(老师)
			v1.POST("modify/teacher/curriculum/video/catalog",v1post.ModifyTeacherCurriculumVideoCatalog)
			// 删除视频目录信息(老师)
			v1.DELETE("del/teacher/curriculum/video/catalog",v1del.DelTeacherCurriculumVideoCatalog)
			// 添加课程的视频(老师)
			v1.POST("add/teacher/curriculum/video/catalog",v1post.AddTeacherCurriculumVideoCatalog)


			// 下架视频
			v1.DELETE("del/curriculum",v1del.DelCurriculum)
			// 查看下架的视频
			v1.GET("show/del/curriculum",v1get.ShowDelCurriculum)
			// 恢复下架视频(老师)
			v1.POST("recovery/curriculum",v1post.RecoveryCurriculum)

			// 查看课程中的视频(目录)  和无需登录的冲突了  解决了 无需登录那个不返回课程的每个url
			v1.GET("see/curriculum/video/:cid",v1get.SeeVideos)

			// 发表评论
			v1.POST("add/comment",v1post.AddComment)
			// 查看评论
			v1.GET("see/comment",v1get.SeeComment)
			// 删除评论
			v1.DELETE("del/comment",v1del.DelComment)

			// 添加购物车,不是直接订单
			v1.POST("add/shopping",v1post.AddShopping)
			// 查看购物车
			v1.GET("show/shopping",v1get.ShowShopping)


			// --------------------------------------------------------------------------------------------

			// 获得微信,支付宝 跳转url(wap版和pc版)  参数 课程id
			v1.GET("get/zfb/:cid",v1get.GetZFBQR)
			// gg 微信开发号给冻结...用不了
			v1.GET("get/wx/:cid",v1get.GetWXQR)

			// todo 由于中间件拦截 所以需要移动到外层
			// 微信 支付宝回调函数  now modify
			// v1.POST("zfb/callback",v1post.ZfbCallback)
			//v1.POST("wx/callback",v1post.WxCallback)
			// 支付宝查询付款详情(是否成功)
			v1.POST("query/zfb/issuccess",v1post.QueryZfbIssuccess)
			v1.POST("query/wx/issuccess",v1post.QueryWxIssuccess)

			// --------------------------------------------------------------------------------------------




			// 购物车下单状态更改(购买) 添加订单
			v1.POST("modify/shopping/status",v1post.ModifyShoppingStatus)

			// 查看课程购买记录(学生or老师)
			v1.GET("show/curriculum/record",v1get.Record)
			// 查看课程被购买记录(老师) 含价格为0的课程
			v1.GET("show/Teacher/curriculum/record",v1get.TeacherRecord)
			// 查看拥有金额(老师)
			v1.GET("have/money",v1get.HaveMoney)
			// 提成金额(老师)  zfb
			v1.POST("extract/money",v1post.ExtractMoney)
			// 提成记录(老师)
			v1.GET("extract/record",v1get.ExtractRecord)

			// 退出
			v1.POST("logout",v1post.Logout)
		}
	}




	// 后台管理,或者使用另一个web服务开后台管理,不过记得模型要一样
	admin := Router.Group("api/form/v1/admin")
	// 登录
	admin.POST("login",adminpost.Login)
	// todo 注册不予开放,直接数据库或者另一个程序进行注册  模型中添加是谁删除的 在admin的删除相关中需要弄好
	admin.POST("register/user",adminpost.RegisterAdmin)
	{
		admin.Use(middleware.AuthAdminLogin())
		// 推出
		admin.POST("logout",adminpost.Logout)

		// 查看视频和总个数
		admin.GET("show/video/total",adminget.ShowVideoTotal)
		// 根据天数查看当天视频上传个数 使用参数控制返回数据范围 1 - 365
		admin.GET("show/curriculum/count/:day",adminget.ShowCurriculumDays)
		// 根据天数查看正常注册的人数(学生or老师or全部)
		admin.GET("show/user/count/:day",adminget.ShowUserDays)

		// 查看总注册人数(老师or学生or全部)
		admin.GET("show/user/total",adminget.ShowUserTotal)
		// 查看用户列表
		admin.GET("show/users/list",adminget.ShowUsersList)

		// 封禁用户 todo 使用用户中的status字段 添加 默认3为封禁用户
		admin.DELETE("del/prohibit/user",admindel.ProhibitUser)
		// 查看封禁用户
		admin.GET("show/prohibit/users",adminget.ShowProhibitUsers)
		// 解封用户
		admin.POST("adopt/user",adminpost.AdoptUser)

		// 查看视频
		admin.GET("show/list/video",adminget.ShowListVideo)
		// 下架视频
		admin.DELETE("del/video",admindel.DelVideo)
		// 恢复视频
		admin.POST("adopt/video",adminpost.AdoptVideo)

		// 查看评论
		admin.GET("show/comments",adminget.ShowComment)
		// 删除评论
		admin.DELETE("del/comment",admindel.DelComment)
		// 恢复评论
		admin.POST("adopt/comment",adminpost.AdoptComment)



		// 根据天数查看当天下单金额
		admin.GET("show/day/money/:day",adminget.ShowDayMoney)
		// 根据月份查看当月下单金额
		admin.GET("show/month/money/:month",adminget.ShowMonthMoney)
		// 总共下单金额
		admin.GET("show/total/money",adminget.ShowTotalMoney)

		// 站点提成金额总数
		admin.GET("show/site/royalty/money",adminget.ShowSiteRoyaltyMoney)
		// 根据天数查看当天提成金额
		admin.GET("show/site/royalty/money/:day",adminget.ShowSiteRoyaltyMoneyDay)


		// 被老师提取出去的金额总数
		admin.GET("show/teacher/royalty/money",adminget.ShowTeacherRoyaltyMoney)
		// 根据天数查看被老师提取出去的金额总数
		admin.GET("show/teacher/royalty/money/:day",adminget.ShowTeacherRoyaltyMoneyDay)
	}

	return Router
}