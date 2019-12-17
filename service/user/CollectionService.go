package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type AddCollectionService struct {
	CID int `json:"cid" form:"cid" binding:"required"`
}


func (service *AddCollectionService)AddCollection(c *gin.Context) *serialize.Response {

	uid := service2.GetUserId(c)

	sql := "insert into use_collections (u_id,c_id) values (?,?)"
	ok := DB.Transaction(DB.DB.Exec(sql,uid,service.CID))
	if !ok {
		return serialize.DBErr("",nil)
	}

	var coll model.UseCollections
	DB.DB.Where("u_id = ?",uid).First(&coll)

	return serialize.Res(nil,"add success")
}
