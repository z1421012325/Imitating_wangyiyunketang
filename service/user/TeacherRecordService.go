package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

type blenddata struct {
	model.Curriculums
	model.ShoppingCarts
	model.User
}

type TeacherRecord struct {
	Result []blenddata						`json:"result"`
	Total  string							`json:"total" gorm:"column:total"`
}

func TeacherRecordService(c *gin.Context)*serialize.Response{

	uid := service.GetUserId(c)
	start,size := service.PagingQuery(c)

	var data TeacherRecord  // 课程 用户 购买记录时的价格
	sql := "select " +
				"u.u_id,u.nickename,u.r_id,cc.c_name,cc.c_image,cc.price " +
			"from " +
				"shopping_carts as sp join curriculums as cc " +
			"on " +
				"sp.c_id = cc.c_id join users as u " +
			"on " +
				"u.u_id = sp.u_id " +
			"where " +
				"sp.c_id in (select c_id from curriculums where u_id = ?) order by sp.create_at desc " +
				"limit ?,?"
	DB.DB.Raw(sql,uid,start,size).Scan(&data.Result)

	countsql := "select count(*) as total from shopping_carts where c_id in (select c_id from curriculums where u_id = ?)"
	DB.DB.Raw(countsql,uid).Scan(&data.Total)

	return serialize.Res(data,"")
}
