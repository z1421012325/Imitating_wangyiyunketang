package admin

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type DelCommentServic struct {
	Cid string		`form:"cid" json:"cid"`
	Uid string		`form:"uid" json:"uid"`
}

func (service *DelCommentServic)DelComment(c *gin.Context)*serialize.Response{

	aid := service2.GetUserId(c)

	sql := "update curriculum_comments set admin_del = now(),a_id = ? where c_id = ? and u_id = ?"

	db1 := DB.DB.Exec(sql,aid,service.Cid,service.Uid)

	ok := DB.Transaction(db1)
	if !ok{
		return serialize.DBErr("delete faild",nil)
	}

	return serialize.Res(nil,"delete success")

}







