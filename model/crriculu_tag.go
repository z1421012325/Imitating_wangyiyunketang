package model

/*
课程与标签 第三方表 多对多
CREATE TABLE `crriculu_tag` (
  `c_id` int(10) DEFAULT NULL COMMENT '外键 课程id 第三方表',
  `t_id` int(11) DEFAULT NULL COMMENT '外键 标签tagsid 第三方表',
  KEY `c_id` (`c_id`),
  KEY `t_id` (`t_id`),
  CONSTRAINT `crriculu_tag_ibfk_1` FOREIGN KEY (`c_id`) REFERENCES `curriculums` (`c_id`),
  CONSTRAINT `crriculu_tag_ibfk_2` FOREIGN KEY (`t_id`) REFERENCES `tags` (`t_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
 */
type CrriculuTag struct {
	CID 	int 	`gorm:"column:c_id"`
	TID 	int 	`gorm:"column:t_id"`
}

func (cl *CrriculuTag)TableName()string{
	return "crriculu_tag"
}