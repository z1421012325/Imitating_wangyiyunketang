package model

import "time"

/*
购物车
CREATE TABLE `shopping_carts` (
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键 用户id',
  `c_id` int(10) DEFAULT NULL COMMENT '外键 课程id',
  `number` int(11) DEFAULT '1' COMMENT '课程数量,但是是网易云课堂类似的,默认就是1买把...',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  KEY `u_id` (`u_id`),
  KEY `c_id` (`c_id`),
  CONSTRAINT `shopping_carts_ibfk_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`),
  CONSTRAINT `shopping_carts_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `curriculums` (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 */
type ShoppingCarts struct {
	UID        int       `gorm:"column:u_id" json:"uid"`
	CID        int       `gorm:"column:c_id" json:"cid"`
	Number     int       `gorm:"column:number" json:"count"`
	CreateTime time.Time `gorm:"column:create_at" json:"at"`
}

func (sc *ShoppingCarts)TableName()string{
	return "shopping_carts"
}
