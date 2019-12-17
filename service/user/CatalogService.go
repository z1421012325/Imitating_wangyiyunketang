package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)



func CatalogService(c *gin.Context) *serialize.Response{
	cid := c.Param("cid")

	var data []model.CataLog
	//DB.DB.Where("c_id = ?",cid).Order("create_at ASC").Find(&data)
	sql := "select c_id,name,create_at from catalog where c_id = ? order by create_at asc"
	DB.DB.Raw(sql,cid).Scan(&data)
	return serialize.Res(data,"")
}
