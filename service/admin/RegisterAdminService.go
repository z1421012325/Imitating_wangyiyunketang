package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
)

type RegisterAdminService struct {
	Username string		 `json:"username" form:"username"   binding:"required,min=2,max=30"`
	Pswd     string		 `json:"pswd"     form:"pswd"       binding:"required,min=2,max=30"`
	Status   int         `json:"status"   form:"status"`
	Info     string      `json:"info"     form:"info"`
}


func (service *RegisterAdminService)RegisterAdmin()*serialize.Response{

	var user model.Admin
	DB.DB.Where("username = ?",service.Username).First(&user)
	if user.UserName == service.Username{
		return serialize.Res(nil,"账号已存在")
	}

	if err := user.SetPassword(service.Pswd);err != nil {
		return serialize.PswdErr("密码加密失败",nil)
	}

	sql := "insert into admins (username,pswd,status,info,create_at) values (?,?,?,?,now())"
	db1 := DB.DB.Exec(sql,service.Username,user.Pswd,service.Status,service.Info)
	ok := DB.Transaction(db1)
	if !ok {
		return serialize.DBErr("创建账号失败",nil)
	}

	var adminuser model.Admin
	DB.DB.Where("username = ?",service.Username).First(&adminuser)
	adminuser.Pswd = ""
	return serialize.Res(adminuser,"注册成功")
}