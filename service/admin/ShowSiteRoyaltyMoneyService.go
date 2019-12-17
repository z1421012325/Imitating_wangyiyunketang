package admin

import (
	"demos/DB"
	"demos/serialize"
)

type ShowSiteRoyaltyMoneyServiceRes struct {
	Money float64			`gorm:"column:money" json:"money"`
}

func ShowSiteRoyaltyMoneyService()*serialize.Response{

	sql := "select sum(t_money)-sum(actual_money)as money from extracts"
	var data ShowSiteRoyaltyMoneyServiceRes
	DB.DB.Raw(sql).Scan(&data)

	return serialize.Res(data,"")
}
