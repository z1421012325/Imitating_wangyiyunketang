package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)



func CatalogService(c *gin.Context) *serialize.Response{
	cid := c.Param("cid")
	var data []model.CataLog
	DB.DB.Where("c_id = ?",cid).Order("create_at ASC").Find(&data)
	return serialize.Res(data,"")
}
