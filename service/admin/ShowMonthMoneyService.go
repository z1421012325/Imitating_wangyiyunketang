package admin

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)
type ShowMonthMoneyServiceDate struct {
	Time   string		`gorm:"column:time"  json:"time"`
	Money  float64		`gorm:"column:money" json:"money"`
	Count  int			`gorm:"column:count" json:"count"`
}

func ShowMonthMoneyService(c *gin.Context)*serialize.Response{
	month := c.Param("month")

	sql := "select date_format(create_at,'%Y-%m')as time,sum(price)as money,count(*)as count from purchases where status != 0 group by time limit 0,?"
	var data []ShowMonthMoneyServiceDate
	DB.DB.Raw(sql,month).Scan(&data)

	return serialize.Res(data,"")

}
