package admin

import (
	"demos/DB"
	"demos/model"
	"demos/serialize"
)

type AdoptUserServic struct {
	UID 	string  `json:"uid" form:"uid" binding:"required"`
}


func (service *AdoptUserServic)AdoptUser()*serialize.Response{

	sql := "update users set status = ? where status = ? and u_id = ?"
	db1 := DB.DB.Exec(sql,model.Active,model.Prohibit,service.UID)

	ok := DB.Transaction(db1)
	if !ok {
		return serialize.DBErr("adopt user faild",nil)
	}

	return serialize.Res(nil,"success adopt user")
}