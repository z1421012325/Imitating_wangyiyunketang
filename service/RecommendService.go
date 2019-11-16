package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)



type RecommendServiceData struct {
	model.Curriculums
}

func RecommendService(c *gin.Context) *serialize.Response{
	cid := c.Param("cid")
	start,end := pagingQuery(c)

	// 为 true 则输出该老师的视频推荐,热度推荐         热度没搞定
	sql := ""
	if c.DefaultQuery("degree","true") == "true"{
		sql = "select * from curriculums as cc where cc.u_id = (select u_id from curriculums where c_id = ?) limit ?,?"
	}else {	 // false
		sql = "select * from curriculums as cc where cc.t_id = (select t_id from curriculums where c_id = ?) limit ?,?"
	}
	var data []RecommendServiceData
	DB.DB.Raw(sql,cid,start,end).Scan(&data)

	return serialize.Res(data,"")

}
