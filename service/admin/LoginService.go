package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	UserName string  	`json:"username" form:"username"  binding:"required"`
	Pswd 	 string		`json:"pswd"     form:"pswd"      binding:"required"`
}


func (service *LoginService)Login(c *gin.Context)*serialize.Response{

	var adminuser model.Admin
	DB.DB.Where("username = ?",service.UserName).First(&adminuser)

	if !adminuser.CheckPassword(service.Pswd){
		return serialize.PswdErr("密码错误",nil)
	}

	service.SetSession(c,adminuser)

	return serialize.Res(adminuser,"登录成功")
}


func (service *LoginService)SetSession(c *gin.Context,admin model.Admin){

	s := sessions.Default(c)
	s.Clear()
	s.Set("admin_id",admin.ID)
	_ = s.Save()

}

