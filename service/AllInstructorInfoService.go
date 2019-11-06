package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type AllInstructor struct {
	Result []model.User `json:"users"`
	Total  int          `json:"total" gorm:"column:total"`
}

func AllInstructorService(c *gin.Context) *serialize.Response{
	var user AllInstructor
	start,end := pagingQuery(c)

	sql := "select * from users where r_id = ? limit ?,?"
	DB.DB.Raw(sql,model.Teacher,start,end).Find(&user.Result)
	DB.DB.Model(&model.User{}).Where("r_id = ?",model.Teacher).Count(&user.Total)

	return serialize.Res(user,"")
}