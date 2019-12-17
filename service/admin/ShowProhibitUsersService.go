package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"github.com/gin-gonic/gin"
)

type ShowProhibitUsers struct {
	Result []model.User		`json:"result"`
	Total  int				`json:"total" gorm:"column:total"`
}


func ShowProhibitUsersService(c *gin.Context)*serialize.Response{

	start ,size := service.PagingQuery(c)
	var data ShowProhibitUsers
	sql := "select * from users where status = ? order by create_at desc limit ?,?"
	DB.DB.Raw(sql,model.Prohibit,start,size).Scan(&data.Result)
	DB.DB.Model(&model.User{}).Where("status = ?",model.Prohibit).Count(&data.Total)

	return serialize.Res(data,"")

}
