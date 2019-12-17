package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

func SeeVideosService(c *gin.Context)*serialize.Response{

	cid := c.Param("cid")

	var data []model.CataLog
	sql := "select " +
				"c.* " +
			"from " +
				"curriculums as cc join catalog as c " +
			"on " +
				"cc.c_id = c.c_id " +
			"where c.c_id = ? and cc.delete_at is null " +
				"order by c.create_at asc"
	DB.DB.Raw(sql,cid).Scan(&data)

	return serialize.Res(data,"")
}
