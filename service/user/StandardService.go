package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)


func StandardService(c *gin.Context)*serialize.Response{

	uid := service.GetUserId(c)
	if uid == nil{
		return serialize.CheckLogin()
	}

	cid := c.Param("cid")
	// 购买记录或者价格为0
	var ck model.ShoppingCarts
	DB.DB.Where("u_id = ? and c_id = ?",uid,cid).First(&ck)

	var ck1 model.Curriculums
	DB.DB.Where("u_id = ?",uid).First(&ck1)

	if ck.CID == 0 && ck1.Price <= 0.0{
		return serialize.DBErr("未购买该课程",nil)
	}


	var data []model.CataLog
	DB.DB.Where("c_id = ?",cid).Order("create_at asc").Find(&data)

	return serialize.Res(data,"")
}
