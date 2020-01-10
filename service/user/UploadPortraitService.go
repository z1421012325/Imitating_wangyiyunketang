package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type UploadPortraitService struct {
	UID int				`json:"uid" form:"uid"    binding:"required"`
	URL string			`json:"url" form:"url"    binding:"required"`
}

func (service *UploadPortraitService)UploadPortrait(c *gin.Context)*serialize.Response  {

	if !service2.CheckUidToUid(service.UID,c){
		return serialize.CheckLogin()
	}

	sql := "update users set portrait = ? where u_id = ?"
	dbs := DB.DB.Exec(sql,service.URL,service.UID)
	if !DB.Transaction(dbs){
		return serialize.DBErr("保存用户头像失败",nil)
	}

	return serialize.Res(nil,"保存用户头像成功")
}
