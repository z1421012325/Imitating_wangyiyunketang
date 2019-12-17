package user

import (
	"demos/DB"
	"demos/serialize"
	service2 "demos/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DelCollectionService struct {
	CID int `json:"cid" form:"cid" binding:"required"`
}


func (service *DelCollectionService)DelCollection(c *gin.Context) *serialize.Response{
	uid := service2.GetUserId(c)

	sql := "update use_collections set delete_at = now() where u_id = ? and c_id = ?"
	ok := DB.Transaction(DB.DB.Exec(sql,uid,service.CID))
	if !ok{
		return serialize.DBErr("",nil)
	}

	fmt.Println(uid,service.CID)

	return serialize.Res(nil,"del collections success")
}