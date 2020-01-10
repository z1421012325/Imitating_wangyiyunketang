package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)


type Record struct {
	Result []model.Curriculums		`json:"result"`
	Total  int						`json:"total"`
}

func RecordService(c *gin.Context)*serialize.Response{
	uid := service.GetUserId(c)
	Start,size := service.PagingQuery(c)

	sql := "select * from shopping_carts as sp join curriculums as cc on cc.c_id = sp.c_id where sp.u_id = ?  order by sp.create_at desc limit ?,?"
	var data Record
	DB.DB.Raw(sql,uid,Start,size).Scan(&data.Result)
	DB.DB.Model(&model.ShoppingCarts{}).Where("u_id = ?",uid).Count(&data.Total)

	for _,data := range data.Result{
		data.CompletionToOssUrl()
	}

	return serialize.Res(data,"")
}
