package admin

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type ShowDayMoneyServiceRes struct {
	Time   string		`gorm:"column:time" json:"time"`
	Money  float64		`gorm:"column:money" json:"money"`
	Count  int			`gorm:"column:count" json:"count"`
}

func ShowDayMoneyService(c *gin.Context) *serialize.Response  {
	day := c.Param("day")

	sql := "select date_format(create_at,'%Y-%m-%d')as time,sum(price)as money,count(*)as count from purchases where status != 0 group by time order by time desc limit 0,?"
	var data []ShowDayMoneyServiceRes
	DB.DB.Raw(sql,day).Scan(&data)

	return serialize.Res(data,"")
}
