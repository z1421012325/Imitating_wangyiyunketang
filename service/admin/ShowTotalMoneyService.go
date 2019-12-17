package admin

import (
	"demos/DB"
	"demos/serialize"
)

type ShowTotalMoneyServiceRes struct {
	Moneys float64			`gorm:"column:moneys" json:"moneys"`
}

func ShowTotalMoneyService()*serialize.Response{

	sql := "select sum(price)as moneys from purchases where status != 0"
	var data ShowTotalMoneyServiceRes
	DB.DB.Raw(sql).Scan(&data)

	return serialize.Res(data,"")
}
