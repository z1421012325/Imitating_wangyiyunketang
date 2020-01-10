package admin

import (
	"demos/DB"
	"demos/serialize"
)


type ShowTeacherRoyaltyMoneyServiceDate struct {
	Money float64			`gorm:"column:money" json:"money"`
}

func ShowTeacherRoyaltyMoneyService()*serialize.Response{

	sql := "select sum(actual_money)as money from extracts"
	var data ShowTeacherRoyaltyMoneyServiceDate
	DB.DB.Raw(sql).Scan(&data)

	return serialize.Res(data,"")
}
