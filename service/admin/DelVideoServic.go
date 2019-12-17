package admin

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type DelVideoServic struct {
	CID int    `json:"cid" form:"cid"`
}


func (service *DelVideoServic)DelVideo(c *gin.Context)*serialize.Response{

	aid := service2.GetUserId(c)

	sql := "update curriculums set admin_del = now(),a_id = ? where c_id = ? and delete_at is null"
	db1 := DB.DB.Exec(sql,aid,service.CID)
	ok  := DB.Transaction(db1)
	if !ok {
		return serialize.DBErr("del curriculums faild",nil)
	}
	return serialize.Res(nil,"success curriculums")

}