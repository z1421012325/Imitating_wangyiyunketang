package service

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)



// 设计表是没有想那么多,要么在新建一个用户详细信息表,要么用户表user添加一个字段,不想弄,所以只改昵称吧..
type ModifyService struct {
	Nickename 	string		`json:"nickename" form:"nickename" binding:"required,min=5,max=30"`
}



func (service *ModifyService)Modify(c *gin.Context) *serialize.Response {
	uid := GetUserId(c)

	var user model.User
	DB.DB.Where("u_id = ?",uid).First(&user)

	// 修改昵称那个接口测试一下 mysql中开启begin,但是不提交,注意接口发出的请求和mysql中开启事务的请求不是同一个,测试是否阻塞
	ok := DB.Transaction(DB.DB.Model(&user).Where("u_id = ?",uid).
		Update(map[string]interface{}{"nickename":service.Nickename,"status":1}))

	if !ok{
		return serialize.Res(nil,"fail")
	}
	return serialize.Res(nil,"success")
}
