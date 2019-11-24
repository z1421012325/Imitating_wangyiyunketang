package service

import (
	"demos/DB"
	"demos/serialize"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DelCurriculumService struct {
	CID 	int  	`json:"cid" form:"cid" binding:"required"`
}

func (service *DelCurriculumService)DelCurriculum(c *gin.Context)*serialize.Response{

	uid := GetUserId(c)
	sql := "UPDATE curriculums SET delete_at = now() WHERE u_id = ? and c_id = ?"

	ok := DB.Transaction(DB.DB.Exec(sql,uid,service.CID))
	if !ok {
		return serialize.DBErr("",nil)
	}

	fmt.Println(uid,service.CID,sql)
	return serialize.Res(nil,"del curriculum success")
}
