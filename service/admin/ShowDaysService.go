package admin

import (
	"demos/DB"
	"demos/serialize"
	"github.com/gin-gonic/gin"
)

type ress struct {
	Time string    `gorm:"column:time" json:"time"`
	Count int		`gorm:"column:count" json:"count"`
}

func ShowDaysService(c *gin.Context)*serialize.Response  {

	day := c.Param("day")

	sql := "SELECT DATE_FORMAT(create_at,'%Y-%m-%d') as time,count(*)as count FROM curriculums GROUP BY time order by time desc limit 0,?"
	var data []ress
	DB.DB.Raw(sql,day).Scan(&data)

	return serialize.Res(data,"")

}
