package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type NowCurriculumData struct {
	Result []struct{
		model.Curriculums
		// model.CurriculumComment
		Avg string 		`grom:"avg" json:"avg"`
	}
	Total int
}

func NowCurriculumService(c *gin.Context) *serialize.Response{
	uid := c.Param("uid")
	start,end := pagingQuery(c)

	sql := " select " +
				"c.c_id,c.u_id,c.c_name,c.price,c.create_at,avg(cc.number) as avg " +
			"from " +
				"curriculum_comments as cc join curriculums as c " +
			"on " +
				"c.c_id = cc.c_id " +
			"where " +
				"c.u_id = ? and c.delete_at is null " +
			"group by " +
				"c.c_name " +
			"limit " +
				"?,?"

	var now NowCurriculumData
	DB.DB.Raw(sql,uid,start,end).Scan(&now.Result)
	DB.DB.Model(&model.Curriculums{}).Where("u_id = ? and delete_at is null",uid).Count(&now.Total)

	return serialize.Res(now,"")
}
