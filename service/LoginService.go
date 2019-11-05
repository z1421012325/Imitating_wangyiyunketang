package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	Username string		`form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string		`form:"password" json:"password" binding:"required,min=8,max=40"`
}


func (service *LoginService)Login (c *gin.Context) *serialize.Response{

	var user model.User
	DB.DB.Where("username = ?",service.Username).First(&user)
	if user.ID == 0{
		return serialize.AccountErr("",nil)
	}

	if !user.CheckPassword(service.Password){
		return serialize.AccountErr("",nil)
	}


	service.SetSession(c,user)
	return &serialize.Response{
		Msg:  "登录成功",
		Data: &user,
	}

}


func (service *LoginService )SetSession(c *gin.Context,user model.User){
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id",user.ID)
	s.Set("query_control",1)
	_ = s.Save()
}
