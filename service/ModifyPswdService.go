package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ModifyPswdService struct {
	BeforePswd		 string	`json:"beforepswd" form:"beforepswd" binding:"required,min=8,max=40"`
	Pswd             string `json:"pswd" form:"pswd" binding:"required,min=8,max=40"`
	Password_confirm string `json:"pswdconfirm" form:"pswdconfirm" binding:"required,min=8,max=40"`
}

// todo  密码修改,暂时先一个简单的修改
func (service *ModifyPswdService)ModifyPswd(c *gin.Context)*serialize.Response{

	var user model.User
	uid := GetUserId(c)

	ok := service.checkpswd(&user,uid)
	if ok != nil{
		return ok
	}

	// 事务开启
	err := DB.Transaction(DB.DB.Model(&user).
		Where("u_id = ?",uid).Update("pswd",user.Pswd))
	if !err {
		return serialize.DBErr("fail",nil)
	}



	clearSession(c)
	return serialize.Res(nil,"success")
}


// verify密码
func (service *ModifyPswdService) checkpswd(user *model.User,uid interface{}) *serialize.Response{
	if service.Pswd != service.Password_confirm{
		return serialize.PswdErr("",nil)
	}

	DB.DB.Where("u_id = ?",uid).First(&user)

	// 检测原密码和加密之后的密码是否一致
	if !user.CheckPassword(service.BeforePswd){
		return serialize.EncryptionErr("",nil)
	}

	// 加密
	err := user.SetPassword(service.Pswd)
	if err != nil {
		return serialize.EncryptionErr("",nil)
	}

	return nil
}




// 清除当前currnet用户的sesion
func clearSession(c *gin.Context){
	s := sessions.Default(c)
	s.Clear()
	_ = s.Save()
}
