package model

import (
	"demos/DB"
	"demos/util"
	"golang.org/x/crypto/bcrypt"
	"time"
)


/*
用户模型
CREATE TABLE `users` (
  `u_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `nickename` varchar(20) DEFAULT NULL COMMENT '用户昵称',
  `username` varchar(30) NOT NULL COMMENT '用户账号',
  `pswd` varchar(255) DEFAULT NULL COMMENT '密码',
  `status` tinyint(1) DEFAULT '0' COMMENT '外键 用户注册状态,默认为0 未激活',
  `r_id` int(11) DEFAULT '2' COMMENT '身份默认为2,0管理员,1老师 2xuesheng',
  `portrait` varchar(600) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '用户头像 随意设置默认把,如果有统一请设置为这个',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`u_id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `nickename` (`nickename`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
 */
type User struct {
	ID 			uint64		`gorm:"column:u_id" json:"uid"`
	Nickename 	string		`gorm:"column:nickename" json:"nick"`
	Username 	string		`gorm:"column:username" json:"uname"`
	Pswd 		string		`gorm:"column:pswd" json:"pswd"`
	Status 		int			`gorm:"column:status" json:"st"`
	RID 		int			`gorm:"column:r_id" json:"rid"`
	Portrait 	string		`gorm:"column:portrait" json:"img"`
	CreateTime 	time.Time	`gorm:"column:create_at" json:"at"`

	AdminDelTime  	time.Time		`gorm:"column:admin_del" json:"a_del"`		// 后台人员删除时间
	Aid       		int				`gorm:"column:a_id" json:"aid"`				// 后台执行人信息
}

func (u *User)TableName()string{
	return "users"
}


const (
	// PassWordCost 密码加密难度
	PassWordCost = 12

	// 老师
	Teacher int = 1
	// 学生
	Student int = 2

	// 未激活用户
	Inactive = 0
	// 激活用户
	Active = 1
	// 封禁用户
	Prohibit = 3

	// Inactive 未激活用户
	//Inactive string = "inactive"
	// Active 激活用户
	//Active string = "active"

)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.DB.Where("u_id = ?",ID).First(&user)
	user.Pswd = ""
	return user, result.Error
}



// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Pswd = string(bytes)
	return nil
}


// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Pswd), []byte(password))
	return err == nil
}


// 补全存储在oss中的路径
func (model *User)CompletionToOssUrl(){
	model.Portrait = util.CompletionToOssUrl(model.Portrait)
}
