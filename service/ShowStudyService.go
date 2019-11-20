package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"fmt"
	"github.com/gin-gonic/gin"

)

func ShowStudyService(c *gin.Context) *serialize.Response {

	uid := GetUserId(c)
	start,size := pagingQuery(c)

	uid = 100002

	sql := "select " +
		"c.u_id,c.c_id,c.c_name,c.price,c.c_image,c.create_at " +
		"from " +
		"curriculums as c join shopping_carts as sp on c.c_id = sp.c_id " +
		"where " +
		"sp.u_id = ? limit ?,?"

	var data []model.Curriculums
	DB.DB.Raw(sql,uid,start,size).Scan(&data)

	fmt.Println("\n uid 为",uid)

	return serialize.Res(data,"")
}
