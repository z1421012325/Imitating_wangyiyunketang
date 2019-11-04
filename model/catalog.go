package model

import "time"

/*
课程目录
 CREATE TABLE `catalog` (
  `c_id` int(10) DEFAULT NULL COMMENT '外键 课程id',
  `name` varchar(50) DEFAULT NULL COMMENT '课程目录',
  `url` varchar(255) DEFAULT NULL COMMENT '目录地址',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '一个课程多个目录,根据时间排序',
  KEY `c_id` (`c_id`),
  CONSTRAINT `catalog_ibfk_1` FOREIGN KEY (`c_id`) REFERENCES `curriculums` (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
 */
type CataLog struct {
	ID         int       `gorm:"column:c_id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	URL        string    `gorm:"column:url" json:"url"`
	CreateTime time.Time `gorm:"column:create_at" json:"at"`
}

func (cl *CataLog)TableName()string{
	return "catalog"
}