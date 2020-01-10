package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
	"os"
)


type introductionData struct {
	model.User					`json:"user"`
	model.Curriculums			`json:"kc"`
	model.ShoppingCarts			`json:"sp"`		// 根据商品表返回数据确定是否购买
}


func IntroductionService(c *gin.Context) *serialize.Response{
	cid := c.Param("cid")
	if cid == "" {
		return serialize.ParamErr("",nil)
	}

	uid := service.GetUserId(c)

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

	introductionData.User.CompletionToOssUrl()
	introductionData.Curriculums.CompletionToOssUrl()

	// 在redis  zset数据类型中缓存,如果没有则创建一个k为cid的字段,v为1  做热门排行
	ordername := os.Getenv("CLICK_NAME")
	service.SetOrderIncr(ordername,cid)

	return serialize.Res(introductionData,"")
}



