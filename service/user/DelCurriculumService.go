package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type DelCurriculumService struct {
	CID 	int  	`json:"cid" form:"cid" binding:"required"`
}

func (service *DelCurriculumService)DelCurriculum(c *gin.Context)*serialize.Response{

	uid := service2.GetUserId(c)
	sql := "UPDATE curriculums SET delete_at = now() WHERE u_id = ? and c_id = ?"

	ok := DB.Transaction(DB.DB.Exec(sql,uid,service.CID))
	if !ok {
		return serialize.DBErr("del curriculum faild",nil)
	}

	return serialize.Res(nil,"del curriculum success")
}
