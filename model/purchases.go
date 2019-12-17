package model

import "time"

/*
订单
CREATE TABLE `purchases` (
  `c_id` int(10) DEFAULT NULL COMMENT '外键 课程id',
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键 用户id',
  `status` tinyint(1) DEFAULT NULL COMMENT '订单状态,默认为0未支付 支付为1',
  `price` float(10,2) DEFAULT NULL COMMENT '订单当时价格,数量不考虑因为是类似网易云课堂这种 只能买一份',
  `number` varchar(40) DEFAULT NULL COMMENT 'uuid流水号',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  KEY `u_id` (`u_id`),
  KEY `c_id` (`c_id`),
  CONSTRAINT `purchases_ibfk_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`),
  CONSTRAINT `purchases_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `curriculums` (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 */
type Purchases struct {
	CID        int       `gorm:"column:c_id" json:"cid"`
	UID        int       `gorm:"column:u_id" json:"uid"`
	Status     int       `gorm:"column:status" json:"status"`
	Price      float64   `gorm:"column:price" json:"price"`
	Number     string    `gorm:"column:number" json:"count"`
	CreateTime time.Time `gorm:"column:create_at" json:"at"`
}



func (pc *Purchases)TableName()string{
	return "purchases"
}


const (
	DefaultStatus = 0			// 未支付
	CompleteStatus = 1			// 支付
)