package user

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
	"demos/service"
	"os"

	"github.com/gin-gonic/gin"
)

func ShowListClickService(c *gin.Context) *serialize.Response{

	from,to := service.ClickRankingPagingQuery(c)

	ordername := os.Getenv("CLICK_NAME")
	cids := service.GetCacheTypeAggScope(ordername,from,to)

	if data := service.GetCacheTypeStr(cids);data != ""{
		return serialize.Res(data,"")
	}

	var data []model.Curriculums
	DB.DB.Where("c_id in (?) and delete_at is null and admin_del is null",cids).Find(&data)

	service.SetCacheTypeStr(cids,data,0)

	return serialize.Res(data,"")
}
