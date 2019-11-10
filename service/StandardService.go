package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)


type StandardServiceData struct {
	model.Curriculums
	model.ShoppingCarts
}

func StandardService(c *gin.Context)*serialize.Response{

	//user,ok := VerifyUser(c)
	//if !ok{
	//	return serialize.CheckLogin()
	//}
	//user.ID = 100003
	id := 100003
	cid := c.Param("cid")

	sql := "select " +
				"c.c_id,c.u_id,c.c_name,c.price,sc.number " +
			"from " +
				"curriculums as c join shopping_carts as sc " +
			"on " +
				"c.c_id = sc.c_id " +
			"where " +
				"sc.u_id = ? and sc.c_id = ?"
	var check StandardServiceData
	//DB.DB.Raw(sql,user.ID,cid).Scan(&data)
	DB.DB.Raw(sql,id,cid).Scan(&check)

	if check.Curriculums.Price != 0 && check.ShoppingCarts.Number != 1{
		return serialize.Res(nil,"未购买该课程")
	}

	data := CatalogService(c)
	return serialize.Res(data,"")

}
