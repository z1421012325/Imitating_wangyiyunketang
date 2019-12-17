package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)


//type StandardServiceData struct {
//	model.Curriculums
//	model.ShoppingCarts
//}

func StandardService(c *gin.Context)*serialize.Response{

	//id := 100003
	//cid := c.Param("cid")
	//
	//sql := "select " +
	//			"c.c_id,c.u_id,c.c_name,c.price,sc.number " +
	//		"from " +
	//			"curriculums as c join shopping_carts as sc " +
	//		"on " +
	//			"c.c_id = sc.c_id " +
	//		"where " +
	//			"sc.u_id = ? and sc.c_id = ?"
	//var check StandardServiceData
	//DB.DB.Raw(sql,id,cid).Scan(&check)
	//
	//if check.Curriculums.Price != 0 && check.ShoppingCarts.Number != 1{
	//	return serialize.Res(nil,"未购买该课程")
	//}

	//data := CatalogService(c)
	//return serialize.Res(data,"")

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
