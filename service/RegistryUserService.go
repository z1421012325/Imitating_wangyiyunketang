package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"time"
)

type RegistryUserService struct {
	Nickname      string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Username      string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password      string `form:"password" json:"password" binding:"required,min=8,max=40"`
	AgainPassword string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
	Status        int    `form:"status" json:"status"`
}



func (service *RegistryUserService) Register() *serialize.Response{

	if ok := service.Valid(); ok!= nil{
		return ok
	}

	var user model.User
	if err := user.SetPassword(service.Password);err!= nil{
		return serialize.EncryptionErr("",err)
	}


	// gorm 的save 默认就是一个事务
	user.Nickename = service.Nickname
	user.Username  = service.Username
	user.CreateTime = time.Now()
	user.Status    = service.Status

	DB.DB.Save(&user)


	DB.DB.Where("username = ?",service.Username).First(&user)

	res := serialize.Response{
		Code:  0,
		Msg:   "注册成功",
		Data:  &user,
	}
	return &res

}






func (service *RegistryUserService)Valid() *serialize.Response{

	if service.Password != service.AgainPassword {
		return serialize.PswdErr("",nil)
	}


	if service.Status != model.Student && service.Status != model.Teacher {
		return serialize.ParamErr("",nil)
	}


	var count int
	DB.DB.Model(&model.User{}).
		Where("nickename = ? or username = ?",service.Nickname,service.Username).
		Count(&count)
	if count >0 {
		return serialize.AccountErr("",nil)
	}

	return nil

}
