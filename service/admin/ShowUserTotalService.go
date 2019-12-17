package admin

import (
	"demos/DB"
	"demos/serialize"
)

type ShowUserTotalRes struct {
	Count int		`gorm:"column:count" json:"count"`
}

// todo 如果可以尝试查询有效用户,未激活用户,封禁用户

func ShowUserTotalService()*serialize.Response{

	var count ShowUserTotalRes
	//DB.DB.Model(model.User{}).Where("status = ? and status = ?",model.Teacher,model.Student).Count(&Effectivecount)

	sql := "select count(*) as count from users where status != 0"
	DB.DB.Raw(sql).Scan(&count)

	return serialize.Res(count,"")

}
