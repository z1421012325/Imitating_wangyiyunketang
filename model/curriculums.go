package model

import "time"
/*
课程
curriculums | CREATE TABLE `curriculums` (
  `c_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '课程id',
  `c_name` varchar(60) NOT NULL COMMENT '课程名字',
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键 课程老师id',
  `t_id` int(11) DEFAULT NULL COMMENT '外键 tagid',
  `price` float(10,2) DEFAULT '0.00' COMMENT '价格',
  `info` text COMMENT '课程介绍',
  `c_image` varchar(250) DEFAULT NULL COMMENT '阿里云oos直传',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`c_id`),
  KEY `u_id` (`u_id`),
  KEY `t_id` (`t_id`),
  CONSTRAINT `curriculums_ibfk_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`),
  CONSTRAINT `curriculums_ibfk_2` FOREIGN KEY (`t_id`) REFERENCES `tags` (`t_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;
 */
type Curriculums struct {
	CID 		int			`gorm:"column:c_id" json:"cid"`
	UID 		int			`gorm:"column:u_id" json:"uid"`
	//TID 		int			`gorm:"column:t_id" json:"tid"`
	Name 		string		`gorm:"column:c_name" json:"name"`
	Price 		float64		`gorm:"column:price" json:"price"`
	Info		string		`gorm:"column:info" json:"info"`
	Image		string		`gorm:"column:c_image" json:"img"`
	CreateTime 	time.Time	`gorm:"column:create_at" json:"at"`
	DeleteTime 	time.Time	`gorm:"column:delete_at" json:"et"`
}


func (cr *Curriculums)TableName()string{
	return "curriculums"
}


const (
	DefaultPrice  = 0.00
)