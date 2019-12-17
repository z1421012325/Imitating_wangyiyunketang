package model

/*
用户金额
CREATE TABLE `money` (
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键用户表id,唯一',
  `money` float(10,2) DEFAULT '0.00' COMMENT '金钱',
  `version` int(11) DEFAULT NULL COMMENT '乐观锁,版本控制',
  UNIQUE KEY `u_id` (`u_id`),
  CONSTRAINT `money_ibfk_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 */
type Money struct {
	UID     int     `gorm:"column:u_id" json:"uid"`
	Money   float64 `gorm:"column:money" json:"money"`
	Version int     `gorm:"column:version" json:"version"`
}

func (m *Money)TableName()string{
	return "money"
}

const (
	// 划分金额比例
	Divide = 0.05
)
