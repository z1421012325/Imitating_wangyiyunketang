package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type ProhibitUserService struct {
	UID int  `json:"uid" form:"uid"`
}

func (service *ProhibitUserService)ProhibitUser(c *gin.Context)*serialize.Response{

	aid := service2.GetAdminId(c)

	sql := "update users set status = ?,admin_del = now(),a_id = ? where u_id = ? and status != 3"
	db1 := DB.DB.Exec(sql,model.Prohibit,aid,service.UID)
	ok := DB.Transaction(db1)
	if !ok {
		return serialize.DBErr("del user prohibit faild",nil)
	}
	return serialize.Res(nil,"del user prohibit success")
}
