package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

func TeacherCurriculumVideoService(c *gin.Context)*serialize.Response{
	cid := c.Param("cid")
	uid := service.GetUserId(c)

	var data model.CataLog
	sql := "select " +
				"c.* " +
			"from " +
				"users as u join curriculums as cc " +
			"on " +
				"u.u_id = cc.u_id join catalog as c " +
			"on " +
				"c.c_id = cc.c_id " +
			"where u.r_id = ? and c.c_id = ? and c.delete_at is null and cc.u_id = ? " +
				"order by c.create_at asc"
	DB.DB.Raw(sql,model.Teacher,cid,uid).Scan(&data)
	return serialize.Res(data,"")
}
