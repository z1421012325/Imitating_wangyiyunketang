package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)



type RecommendServiceData struct {
	model.Curriculums			`json:"ck"`
}


func RecommendService(c *gin.Context) *serialize.Response{
	cid := c.Param("cid")
	start,end := service.PagingQuery(c)

	// 为 true 则输出该老师的视频推荐,热度推荐
	sql := ""
	isresult := c.DefaultQuery("hot","true")
	if isresult == "true"{
		sql = "select " +
					"* " +
				"from " +
					"curriculums as cc " +
				"where " +
					"cc.u_id = (select u_id from curriculums where c_id = ?) " +
				"limit " +
					"?,?"
	}else {	 // false
		sql = "select " +
					"* " +
				"from " +
					"curriculums as cc " +
				"where " +
					"cc.t_id = (select t_id from curriculums where c_id = ?) " +
				"limit " +
					"?,?"
	}


	// redis 查询
	keys := []string{cid,isresult}
	cachedata := service.GetCacheTypeStr(keys)
	if cachedata != "" {
		return serialize.Res(cachedata,"")
	}


	var data []RecommendServiceData
	DB.DB.Raw(sql,cid,start,end).Scan(&data)

	for _,data := range data {
		data.CompletionToOssUrl()
	}

	// redis 缓存
	service.SetCacheTypeStr(keys,data,0)

	return serialize.Res(data,"")

}
