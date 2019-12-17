package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

type ExtractRecord struct {
	Result 	[]model.Extracts  	`json:"result"`
	Total 	int    				`json:"total"`
}

func ExtractRecordService(c *gin.Context)*serialize.Response{

	uid := service.GetUserId(c)
	start,size := service.PagingQuery(c)

	var data ExtractRecord
	sql := "select * from extracts where u_id = ? order by create_at desc limit ?,?"
	DB.DB.Raw(sql,uid,start,size).Scan(&data.Result)

	DB.DB.Model(&model.Extracts{}).Where("u_id = ?",uid).Count(&data.Total)

	return serialize.Res(data,"")
}
