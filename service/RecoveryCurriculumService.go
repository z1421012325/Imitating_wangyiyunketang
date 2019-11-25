package service

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type RecoveryCurriculumService struct {
	CID 	int  	`json:"cid" form:"cid" binding:"required"`
}


func (service *RecoveryCurriculumService)RecoveryCurriculum(c *gin.Context) *serialize.Response{
	uid := GetUserId(c)

	sql := "update curriculums set delete_at = null where u_id = ? and c_id = ?"
	ok := DB.Transaction(DB.DB.Exec(sql,uid,service.CID))
	if !ok{
		return serialize.DBErr("",nil)
	}

	return serialize.Res(nil,"recovery success")
}