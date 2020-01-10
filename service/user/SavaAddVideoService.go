package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	service2 "demos/service"
	"github.com/gin-gonic/gin"
)

type SavaAddVideoService struct {
	UID    int		`json:"uid"    form:"uid"   binding:"required"`
	CID    int		`json:"cid"    form:"cid"   binding:"required"`
	Name   string   `json:"name"   form:"name"  binding:"required"`
	URL    string	`json:"url"    form:"url"   binding:"required"`
}

func (service *SavaAddVideoService)SavaAddVideo(c *gin.Context) *serialize.Response {

	//uid := service2.GetUserId(c)
	//if uid != uint64(service.UID){
	//	return serialize.CheckLogin()
	//}

	if !service2.CheckUidToUid(service.UID,c){
		return serialize.CheckLogin()
	}


	var c1 model.CataLog
	sql1 := "select * from catalog as c join curriculums as cc on c.c_id = cc.c_id where cc.c_id = ? and cc.u_id = ? limit 0,1"
	DB.DB.Raw(sql1,service.CID,service.UID).Scan(&c1)
	if c1.ID == 0 {
		return serialize.Res(nil,"cid is null")
	}

	sql2 := "insert into catalog (c_id,name,url) values (?,?,?)"
	db1 := DB.DB.Exec(sql2,service.CID,service.Name,service.URL)
	if !DB.Transaction(db1){
		return serialize.DBErr("目录视频保存失败",nil)
	}

	var c2 model.CataLog
	sql3 := "select * from catalog where c_id = ? order by create_at desc limit 0,1"
	DB.DB.Raw(sql3,service.CID).Scan(&c2)

	return serialize.Res(c2,"目录视频保存成功")
}
