package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SeeComment struct {
	Result []model.CurriculumComment		`json:"result"`
	Total int							`json:"total"`
}

func SeeCommentService(c *gin.Context)*serialize.Response{
	uid := GetUserId(c)
	start,size := pagingQuery(c)

	var data SeeComment
	sql := "select * from curriculum_comments where u_id = ? and delete_at is null order by create_at desc limit ?,?"
	DB.DB.Raw(sql,uid,start,size).Scan(&data.Result)
	DB.DB.Model(&model.CurriculumComment{}).Where("u_id = ? and delete_at is null",uid).Count(&data.Total)

	fmt.Println(uid,start,size)
	return serialize.Res(data,"")
}
