package admin

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type ShowuserDaysRes struct {
	Time string  `gorm:"column:time"  json:"time"`
	Count int    `gorm:"column:count" json:"count"`
}


func ShowUserDaysService(c *gin.Context)*serialize.Response{

	day := c.Param("day")

	sql := "select date_format(create_at,'%Y-%m-%d')as time,count(*)as count from users where status != 0 group by time order by time desc limit 0,?"
	var data []ShowuserDaysRes
	DB.DB.Raw(sql,day).Scan(&data)

	return serialize.Res(data,"")

}
