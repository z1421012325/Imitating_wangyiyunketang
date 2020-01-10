package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type ModifyShoppingStatusService struct {
	Number     string    `form:"number" json:"number" binding:"required"`   // 订单流水号
	//UID        int       `form:"uid" json:"uid"`
	//CID        int       `form:"cid" json:"cid"`
}

func (service *ModifyShoppingStatusService)ModifyShoppingStatus(c *gin.Context)*serialize.Response{

	//todo

	var data model.Purchases
	DB.DB.Where("number = ? and status = ?",service.Number,model.DefaultStatus).First(&data)
	if data.Number == ""{
		return serialize.QueryErr("not in uuid",nil)
	}

	sql1 := "update purchases set status = ? where number = ? and status = ?"
	sql2 := "insert into shopping_carts (c_id,u_id) values (?,?)"
	ok := DB.Transaction(DB.DB.Exec(sql1,model.CompleteStatus,service.Number,model.DefaultStatus),
		DB.DB.Exec(sql2,data.CID,data.UID))
	if !ok{
		return serialize.DBErr("add faild",nil)
	}

	return serialize.Res(data,"update success")
}