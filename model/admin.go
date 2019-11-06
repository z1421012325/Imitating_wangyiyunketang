package model

import "time"

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
	ID          int       `gorm:"column:c_id" json:"aid"`
	UserName    string    `gorm:"username:c_id" json:"uname"`
	Pswd        string    `gorm:"pswd:c_id" json:"-"`
	Status      int       `gorm:"column:status" json:"status"`
	Info        string    `gorm:"column:info" json:"info"`
	CcreateTime time.Time `gorm:"column:create_at" json:"at"`
}

func (a *Admin)TableName()string{
	return "admins"
}