package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type AddTeacherCurriculumVideoCatalogService struct {
	ID         int       `json:"cid" form:"cid"`
	Name       string    `json:"name" form:"name"`
	URL        string    `json:"url" form:"url"`

	UID 	   string    `json:"uid" form:"uid"`
}


func (service *AddTeacherCurriculumVideoCatalogService)AddTeacherCurriculumVideoCatalog(c *gin.Context)*serialize.Response{

	uid := service2.GetUserId(c)
	if uid != service.UID{
		return serialize.CheckLogin()
	}

	sql := "insert into catalog (c_id,name,url) values (?,?,?)"
	ok := DB.Transaction(DB.DB.Exec(sql,service.ID,service.Name,service.UID))
	if !ok {
		return serialize.DBErr("add video faild",nil)
	}

	return serialize.Res(nil,"add video success")
}