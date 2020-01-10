package user

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type DelTeacherCurriculumVideoCatalogService struct {
	ID         int       `json:"cid"  form:"cid"`					// 课程id
	CataId     int		 `json:"id"   form:"id"`					// 一个课程中的video的id
}

func (service *DelTeacherCurriculumVideoCatalogService)DelTeacherCurriculumVideoCatalog(c *gin.Context)*serialize.Response{


	sql := "update catalog set delete_at = now() where c_id = ? and id = ?"
	ok := DB.Transaction(DB.DB.Exec(sql,service.ID,service.CataId))
	if !ok{
		return serialize.DBErr("del faild",nil)
	}

	return serialize.Res(nil,"del success")


}
