package admin

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)
type ShowSiteRoyaltyMoneyDayServiceRes struct {
	Time  string			`gorm:"column:time"  json:"time"`
	Money float64			`gorm:"column:money" json:"money"`
}

func ShowSiteRoyaltyMoneyDayService(c *gin.Context)*serialize.Response{
	day := c.Param("day")

	sql := "select date_format(create_at,'%Y-%m-%d')as time,sum(t_money)-sum(actual_money)as money from extracts group by time order by time desc limit 0,?"
	var data []ShowSiteRoyaltyMoneyDayServiceRes
	DB.DB.Raw(sql,day).Scan(&data)

	return serialize.Res(data,"")
}
