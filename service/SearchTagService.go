package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

func SearchTagService(c *gin.Context) *serialize.Response{

	tag   := c.Query("key")
	isall := c.DefaultQuery("all","false")

	sql := ""
	if isall == "true" {
		sql = "select * from tags"
	}else {
		sql = "select * from tags where t_name like '%"+tag+"%'"
	}

	var tags []model.Tag
	DB.DB.Raw(sql).Scan(&tags)

	return serialize.Res(tags,"")
}
