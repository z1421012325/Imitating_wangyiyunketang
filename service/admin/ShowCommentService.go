package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

type ShowCommentRes struct {
	Result   []model.CurriculumComment			`json:"result"`
	Total    string								`json:"total" gorm:"total"`
}


func ShowCommentService(c *gin.Context) *serialize.Response  {

	start,size := service.PagingQuery(c)

	sql1 := "select * from curriculum_comments where admin_del is null and delete_at is null order by create_at desc limit 0,?"

	var data ShowCommentRes
	DB.DB.Raw(sql1,start,size).Scan(&data.Result)

	sql2 := "select count(*)as total from curriculum_comments  where admin_del is null and delete_at is null"
	DB.DB.Raw(sql2).Scan(&data.Total)

	return serialize.Res(data,"")

}
