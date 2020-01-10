package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"github.com/gin-gonic/gin"
	"demos/service"
)

type ShowUsersListRes struct {
	Result 	[]model.User		`json:"result"`
	Total 	string 				`json:"total"`
}

func ShowUsersListService(c *gin.Context)*serialize.Response{

	start,size := service.PagingQuery(c)

	sql := "select u_id,nickename,status,r_id,portrait,create_at from users where status != 0 order by create_at desc limit ?,?"
	var data ShowUsersListRes
	DB.DB.Raw(sql,start,size).Scan(&data.Result)
	DB.DB.Model(&model.User{}).Where("status != ? ",0).Count(&data.Total)

	for _,data := range data.Result{
		data.CompletionToOssUrl()
	}

	return serialize.Res(data,"")

}
