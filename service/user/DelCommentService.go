package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)



type DelCommentService struct {
	CID     int 	`json:"cid" form:"cid" binding:"required"`
}

func (service *DelCommentService)DelComment(c *gin.Context)*serialize.Response{
	uid := service2.GetUserId(c)

	sql := "update curriculum_comments set delete_at = now() where u_id = ? and c_id = ?"
	ok := DB.Transaction(DB.DB.Exec(sql,uid,service.CID))
	if !ok {
		return serialize.DBErr("del faild",nil)
	}

	return serialize.Res(nil,"del success")
}