package user

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type ModifyTeacherCurriculumVideoService struct {
	ID         int       `json:"cid"   form:"cid"    binding:"required"`
	Name       string    `json:"name"  form:"name"   binding:"required"`

	CataId     int		 `json:"id"    form:"id"     binding:"required"`
}


func (service *ModifyTeacherCurriculumVideoService)ModifyTeacherCurriculumVideo(c *gin.Context)*serialize.Response{

	sql := "update cataLog set name = ? where c_id = ? and id = ?"
	ok := DB.Transaction(DB.DB.Exec(sql,service.Name,service.ID,service.CataId))
	if !ok{
		return serialize.DBErr("update faild",nil)
	}
	return serialize.Res(nil,"update success")

}