package model

import "time"
/*
金额提取
CREATE TABLE `extracts` (
  `e_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键自增长',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `t_money` float(10,2) DEFAULT NULL COMMENT '提取金额',
  `divide` float(10,2) DEFAULT NULL COMMENT '站点分成,默认为5%',
  `actual_money` float(10,2) DEFAULT NULL COMMENT '实际提成',
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键 用户id',
  `number` varchar(32) DEFAULT NULL COMMENT '流水号',
  PRIMARY KEY (`e_id`),
  KEY `u_id` (`u_id`),
  CONSTRAINT `extracts_ibfk_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8
 */
type Extracts struct {
	EID         int       `gorm:"column:e_id" json:"eid"`
	UID         int       `gorm:"column:" json:"uid"`
	Money       float64   `gorm:"column:t_money" json:"money"`
	Divide      float64   `gorm:"column:divide" json:"divide"`
	ActualMoney float64   `gorm:"column:actual_money" json:"am"`
	Number      string    `gorm:"column:number" json:"sid"`
	CreateTime  time.Time `gorm:"column:create_at" json:"at"`
}

func (ec *Extracts)TableName()string{
	return "extracts"
}