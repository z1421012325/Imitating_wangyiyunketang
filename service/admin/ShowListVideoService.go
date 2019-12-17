package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

type ShowListVideoRes struct {
	Result []model.Curriculums  `json:"result"`
	Total  int					`json:"total"`
}

func ShowListVideoService(c *gin.Context)*serialize.Response{

	start,size := service.PagingQuery(c)
	sql := "select * from curriculums where delete_at is null order by create_at desc limit ?,?"
	var data ShowListVideoRes
	DB.DB.Raw(sql,start,size).Scan(&data.Result)
	DB.DB.Model(&model.Curriculums{}).Where("delete_at is null").Count(&data.Total)

	return serialize.Res(data,"")
}
