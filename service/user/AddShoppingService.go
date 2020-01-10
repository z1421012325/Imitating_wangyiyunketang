package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	service2 "demos/service"
	"demos/util"
	"github.com/gin-gonic/gin"
)

type AddShoppingService struct {
	CID int `json:"cid" form:"cid" binding:"required"`
}


func (service *AddShoppingService)AddShopping(c *gin.Context)*serialize.Response{

	uid := service2.GetUserId(c)

	var price model.Curriculums
	DB.DB.Select("price").Where("c_id = ? and delete_at is null",service.CID).First(&price)

	uuid := util.GetUuid()

	sql := "insert into purchases (u_id,c_id,status,price,number) values (?,?,?,?,?)"
	db1 := DB.DB.Exec(sql,uid,service.CID,model.DefaultStatus,price.Price,uuid)
	ok := DB.Transaction(db1)
	if !ok {
		return serialize.DBErr("add shopping faild",nil)
	}

	return serialize.Res(nil,"add shopping success")
}
