package admin

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"

	s2 "demos/service"
)

type AdoptCommentService struct {
	CID int			`json:"cid" form:"cid" binding:"required"`
	UID int			`json:"uid" form:"uid" binding:"required"`
}

func (service *AdoptCommentService)AdoptComment(c *gin.Context)*serialize.Response  {

	aid := s2.GetAdminId(c)

	sql := "update curriculum_comments set admin_del = null,a_id = ? where c_id = ? and u_id = ?"
	dbs := DB.DB.Exec(sql,aid,service.CID,service.UID)
	if !DB.Transaction(dbs){
		return serialize.DBErr("update fail comment",nil)
	}

	return serialize.Res(nil,"update success comment")

}
