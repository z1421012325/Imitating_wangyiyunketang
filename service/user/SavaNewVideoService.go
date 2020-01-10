package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SavaNewVideoService struct {
	UID         int			`json:"uid"     form:"uid"      binding:"required"`
	Title 		string		`json:"title"   form:"title"    binding:"required"`
	Price 		float64		`json:"price"   form:"price"`
	Info		string		`json:"info"    form:"info"`
	ImageUrl	string		`json:"i_url"   form:"i_url"`

	VideoName    string		`json:"name"    form:"name"     binding:"required"`
	VideoURL    string		`json:"url"     form:"url"      binding:"required"`
}

func (service *SavaNewVideoService) SavaNewVideo(c *gin.Context) *serialize.Response {

	fmt.Println(service)

	uid := service2.GetUserId(c)
	if uid != uint64(service.UID){
		return serialize.CheckLogin()
	}

	sql1 := "insert into curriculums (u_id,c_name,price,info,c_image,create_at) values (?,?,?,?,?,now())"
	db1 := DB.DB.Exec(sql1,service.UID,service.Title,service.Price,service.Info,service.ImageUrl)
	if !DB.Transaction(db1){
		return serialize.Res(nil,"创建课程失败")
	}


	sql2 := "insert into catalog (c_id,name,url) values ((select c_id from curriculums where u_id = ? and c_name = ?),?,?)"
	db2 := DB.DB.Exec(sql2,service.UID,service.Title,service.VideoName,service.VideoURL)
	if !DB.Transaction(db2){
		return serialize.Res(nil,"课程保存失败")
	}

	return serialize.Res(nil,"创建课程并保存成功")
}