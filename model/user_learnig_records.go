package model

/*
用户学习记录
CREATE TABLE `user_learnig_records` (
  `c_id` int(10) DEFAULT NULL COMMENT '外键 课程id',
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键 用户id  未完成功能记录每个课程每个用户在那个学习目录',
  KEY `c_id` (`c_id`),
  KEY `u_id` (`u_id`),
  CONSTRAINT `user_learnig_records_ibfk_1` FOREIGN KEY (`c_id`) REFERENCES `curriculums` (`c_id`),
  CONSTRAINT `user_learnig_records_ibfk_2` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
 */
type UserLearnigRecords struct {
	CID int			`gorm:"column:c_id"`
	UID int			`gorm:"column:u_id"`
}

func (ulr *UserLearnigRecords)TableName()string{
	return "user_learnig_records"
}
