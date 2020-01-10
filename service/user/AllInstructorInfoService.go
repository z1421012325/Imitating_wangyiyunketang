package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	_"demos/service"
	"github.com/gin-gonic/gin"
)

type AllInstructorData struct {
	Result []model.User `json:"users"`
	Total  int          `json:"total" gorm:"column:total"`
}

func AllInstructorService(c *gin.Context) *serialize.Response{
	var data AllInstructorData
	start,end := service.PagingQuery(c)

	sql := "select * from users where r_id = ? limit ?,?"
	DB.DB.Raw(sql,model.Teacher,start,end).Find(&data.Result)
	DB.DB.Model(&model.User{}).Where("r_id = ?",model.Teacher).Count(&data.Total)

	for _, data := range data.Result{
		data.CompletionToOssUrl()
	}

	return serialize.Res(data,"")
}