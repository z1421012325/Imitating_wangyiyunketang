package service

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type AddCommentService struct {
	CID 		string 		`json:"cid" form:"cid" binging:"required"`
	Msg 		string		`json:"msg" form:"msg" binging:"required"`
	Number   	int       	`json:"number" form:"number"`
}

func (service *AddCommentService)AddComment(c *gin.Context)*serialize.Response{
	uid := GetUserId(c)

	sql := "insert into curriculum_comments (u_id,c_id,number,comment) values (?,?,?,?)"
	ok := DB.Transaction(DB.DB.Exec(sql,uid,service.CID,service.Number,service.Msg))
	if !ok{
		return serialize.DBErr("",nil)
	}

	return serialize.Res(nil,"add comment success")
}