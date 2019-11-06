package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)



func CourseDetailService(c *gin.Context) *serialize.Response{
	cid := c.Param("cid")
	var cl []model.CataLog
	DB.DB.Where("c_id = ?",cid).Order("create_at").Find(&cl)
	return serialize.Res(cl,"")
}