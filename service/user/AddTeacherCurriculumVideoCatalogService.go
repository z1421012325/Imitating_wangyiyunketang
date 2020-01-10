package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type AddTeacherCurriculumVideoCatalogService struct {
	ID         int       `json:"cid"   form:"cid"    binding:"required"`
	Name       string    `json:"name"  form:"name"   binding:"required"`
	URL        string    `json:"url"   form:"url"    binding:"required"`

	UID 	   int       `json:"uid"   form:"uid"    binding:"required"`
}


func (service *AddTeacherCurriculumVideoCatalogService)AddTeacherCurriculumVideoCatalog(c *gin.Context)*serialize.Response{

	if !service2.CheckUidToUid(service.UID,c){
		return serialize.CheckLogin()
	}

	sql := "insert into catalog (c_id,name,url) values (?,?,?)"
	ok := DB.Transaction(DB.DB.Exec(sql,service.ID,service.Name,service.UID))
	if !ok {
		return serialize.DBErr("add video faild",nil)
	}

	return serialize.Res(nil,"add video success")
}