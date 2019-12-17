package admin

import (
	"demos/DB"
	"demos/serialize"
)

type res struct {
	Count int		`json:"count" gorm:"column:count"`
}

func ShowTotalService()*serialize.Response{

	var data res
	sql := "select count(*) as count from curriculums where delete_at is null"
	DB.DB.Raw(sql).Scan(&data)

	return serialize.Res(data,"")
}
