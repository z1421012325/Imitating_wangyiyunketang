package user

import (
	"demos/DB"
	"demos/serialize"
	"fmt"
	"github.com/gin-gonic/gin"

)

type DelTeacherCurriculumVideoCatalogService struct {
	ID         int       `gorm:"column:c_id" json:"cid"`
	CataId     int		 `gorm:"column:id" json:"id"`
}

func (service *DelTeacherCurriculumVideoCatalogService)DelTeacherCurriculumVideoCatalog(c *gin.Context)*serialize.Response{

	fmt.Println(service.ID,service.CataId)

	sql := "update catalog set delete_at = now() where c_id = ? and id = ?"
	ok := DB.Transaction(DB.DB.Exec(sql,service.ID,service.CataId))
	if !ok{
		return serialize.DBErr("del faild",nil)
	}

	return serialize.Res(nil,"del success")


}
