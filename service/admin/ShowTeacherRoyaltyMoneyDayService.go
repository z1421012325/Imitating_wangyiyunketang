package admin

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type ShowTeacherRoyaltyMoneyDayServiceRes struct {
	Time  string			`gorm:"column:time"  json:"time"`
	Money float64			`gorm:"column:money" json:"money"`
}

func ShowTeacherRoyaltyMoneyDayService(c *gin.Context)*serialize.Response{
	day := c.Param("day")

	sql := "select date_format(create_at,'%Y-%m-%d')as time,sum(actual_money)as money from extracts group by time limit 0,?"
	var data []ShowTeacherRoyaltyMoneyDayServiceRes
	DB.DB.Raw(sql,day).Scan(&data)

	return serialize.Res(data,"")
}
