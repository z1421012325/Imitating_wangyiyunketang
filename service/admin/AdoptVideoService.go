package admin

import (
	"demos/DB"
	"demos/serialize"
	"fmt"
	"github.com/gin-gonic/gin"

	s2 "demos/service"
)

type AdoptVideoService struct {
	CID int			`json:"cid" form:"cid" binding:"required"`
}

func (service *AdoptVideoService)AdoptVideo(c *gin.Context)*serialize.Response  {

	aid := s2.GetAdminId(c)
	fmt.Println(aid,service.CID)

	sql := "update curriculums set admin_del = null,a_id = ? where c_id = ?"
	db1 := DB.DB.Exec(sql,aid,service.CID)
	if !DB.Transaction(db1){
		return serialize.DBErr("fail adopt video",nil)
	}
	return serialize.Res(nil,"success adopt video")
}