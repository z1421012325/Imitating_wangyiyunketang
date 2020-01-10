package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)



type NowCurriculumDataSubordinate struct {
	model.Curriculums
	Avg string 		`grom:"avg" json:"avg"`
}

type NowCurriculumData struct {
	Result []NowCurriculumDataSubordinate		`json:"result"`
	Total int									`json:"total"`
}

func NowCurriculumService(c *gin.Context) *serialize.Response{
	uid := c.Param("uid")
	start,end := service.PagingQuery(c)

	sql := " select " +
				"c.c_id,c.u_id,c.c_name,c.price,c.c_image,c.create_at,avg(cc.number) as avg " +
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

	var data NowCurriculumData
	DB.DB.Raw(sql,uid,start,end).Scan(&data.Result)
	DB.DB.Model(&model.Curriculums{}).Where("u_id = ? and delete_at is null",uid).Count(&data.Total)

	for _,data := range data.Result{
		data.CompletionToOssUrl()
	}

	return serialize.Res(data,"")
}
