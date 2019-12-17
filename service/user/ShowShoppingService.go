package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)


type PurchasesANDCurriculums struct {
	model.Purchases
	model.Curriculums
}

type ShowShopping struct {
	Result 	[]PurchasesANDCurriculums		`json:"result"`
	Total 	int								`json:"total"`
}

func ShowShoppingService(c *gin.Context)*serialize.Response{
	uid := service.GetUserId(c)
	start,size := service.PagingQuery(c)

	var data ShowShopping
	//sql := "select * from purchases where u_id = ? order by create_at desc limit ?,?"
	sql := "select * from purchases as p join curriculums as c on c.c_id = p.c_id where p.u_id = ? order by p.create_at desc limit ?,?"
	DB.DB.Raw(sql,uid,start,size).Scan(&data.Result)

	DB.DB.Model(&model.Purchases{}).Where("u_id = ?",uid).Count(&data.Total)

	return serialize.Res(data,"")
}
