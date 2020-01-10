package model

import (
	"demos/DB"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/*
admin
create table admins (
`aid` tinyint primary key auto_increment comment '管理员id',
`username` char(20) not null comment '账号',
`pswd` varchar(255) not null comment '密码',
`status` tinyint default '0' comment '身份信息',
`info` text comment '一些额外的信息',
`create_at` datetime default current_timestamp,
UNIQUE KEY `username` (`username`))
ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
 */
type Admin struct {
	ID          int       `gorm:"column:aid"           json:"aid"`
	UserName    string    `gorm:"column:username"      json:"username"`
	Pswd        string    `gorm:"column:pswd"          json:"-"`
	Status      int       `gorm:"column:status"        json:"status"`
	Info        string    `gorm:"column:info"          json:"info"`
	CreateTime  time.Time `gorm:"column:create_at"     json:"at"`
}

func (a *Admin)TableName()string{
	return "admins"
}

// GetAdminUser
func GetAdminUser(ID interface{}) (Admin, error) {
	var user Admin
	result := DB.DB.Where("aid = ?",ID).First(&user)
	user.Pswd = ""
	return user, result.Error
}



// SetPassword 设置密码
func (user *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Pswd = string(bytes)
	return nil
}


// CheckPassword 校验密码
func (user *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Pswd), []byte(password))
	return err == nil
}
