package model

import "time"

/*
课程评价
CREATE TABLE `curriculum_comments` (
  `c_id` int(10) DEFAULT NULL COMMENT '外键 课程id',
  `u_id` bigint(20) DEFAULT NULL COMMENT '外键 用户id',
  `number` tinyint(10) DEFAULT NULL COMMENT '评价分数',
  `comment` varchar(300) DEFAULT NULL COMMENT '评价',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `delete_at` datetime DEFAULT NULL COMMENT '删除时间',
  UNIQUE KEY `u_id` (`u_id`,`c_id`),
  KEY `curriculum_comments_ibfk_2` (`c_id`),
  CONSTRAINT `curriculum_comments_ibfk_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`),
  CONSTRAINT `curriculum_comments_ibfk_2` FOREIGN KEY (`c_id`) REFERENCES `curriculums` (`c_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
 */
type CurriculumComment struct {
	CID        int       `gorm:"column:c_id" json:"cid"`
	UID        int       `gorm:"column:u_id" json:"uid"`
	Number     int       `gorm:"column:number" json:"number"`
	Comment    string    `gorm:"column:comment" json:"comment"`
	CreateTime time.Time `gorm:"column:create_at" json:"at"`
	DeleteTime time.Time `gorm:"column:delete_at" json:"-"`

	AdminDelTime  	time.Time		`gorm:"column:admin_del" json:"a_del"`		// 后台人员删除时间
	Aid       		int				`gorm:"column:a_id" json:"aid"`				// 后台执行人信息
}


func (cc *CurriculumComment)TableName()string{
	return "curriculum_comments"
}
