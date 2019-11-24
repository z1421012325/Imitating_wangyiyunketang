package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


type introductionData struct {
	model.User					`json:"user"`
	model.Curriculums
	model.ShoppingCarts			`json:"sp"`		// 根据商品表返回数据确定是否购买
}


func IntroductionService(c *gin.Context) *serialize.Response{
	cid := c.Param("cid")
	if cid == "" {
		return serialize.ParamErr("",nil)
	}

	uid := sessions.Default(c).Get("user_id")
	var introductionData introductionData
	if uid == "" {
		sql := "select " +
					"u.u_id,u.nickename,u.status,u.r_id," +
					"cc.c_id,cc.c_name,cc.u_id,cc.t_id,cc.price,cc.info,cc.c_image,cc.create_at " +
				"from " +
					"users as u join curriculums as cc " +
				"on " +
					"u.u_id = cc.u_id " +
				"where " +
					"cc.c_id = ? limit 1"
		DB.DB.Raw(sql,cid).Scan(&introductionData)
	}else {
		sql := "select " +
					"u.u_id,u.nickename,u.status,u.r_id," +
					"cc.c_id,cc.c_name,cc.u_id,cc.t_id,cc.price,cc.info,cc.c_image,cc.create_at," +
					"sp.c_id,sp.u_id " +
				"from " +
					"users as u join curriculums as cc " +
				"on " +
					"u.u_id = cc.u_id " +
				"join " +
					"shopping_carts sp " +
				"on " +
					"cc.c_id = sp.c_id " +
					"where cc.c_id = ? and sp.u_id = ? limit 1"
		DB.DB.Raw(sql,cid,uid).Scan(&introductionData)
	}

	return serialize.Res(introductionData,"")
}
