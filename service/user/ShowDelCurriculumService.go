package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"

)


type ShowDelCurriculum struct {
	Result []model.Curriculums			`json:"result"`
	Total int						`json:"total"`
}

func ShowDelCurriculumService(c *gin.Context)*serialize.Response{
	uid := service.GetUserId(c)
	start,size := service.PagingQuery(c)

	var data ShowDelCurriculum
	sql := "select * from curriculums where u_id = ? and delete_at is not null limit ?,?"
	DB.DB.Raw(sql,uid,start,size).Scan(&data.Result)
	DB.DB.Model(&model.Curriculums{}).Where("u_id = ? and delete_at is not null").Count(&data.Total)

	return serialize.Res(data,"")
}
