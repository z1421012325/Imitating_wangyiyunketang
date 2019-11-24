package model

import "time"

/*
用户收藏课程
CREATE TABLE `use_collections` (
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键 用户id',
  `c_id` int(10) DEFAULT NULL COMMENT '外键 课程id',
  KEY `c_id` (`c_id`),
  KEY `u_id` (`u_id`),
  CONSTRAINT `use_collections_ibfk_1` FOREIGN KEY (`c_id`) REFERENCES `curriculums` (`c_id`),
  CONSTRAINT `use_collections_ibfk_2` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 */
type UseCollections struct {
	UID 		int 		`gorm:"column:u_id"`
	CID 		int 		`gorm:"column:c_id"`
	Create_At 	time.Time	`gorm:"column:create_at"`
	Delete_At   time.Time   `gorm:"column:delete_at"`
}

func (uc *UseCollections)TableName()string{
	return "use_collections"
}
