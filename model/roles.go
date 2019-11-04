package model

/*
角色
CREATE TABLE `roles` (
  `r_id` tinyint(1) NOT NULL COMMENT '角色id',
  `r_name` char(10) DEFAULT NULL COMMENT '角色名字',
  PRIMARY KEY (`r_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
 */
type Role struct {
	RID  int    `gorm:"column:r_id" json:"rid"`
	Name string `gorm:"column:r_name" json:"name"`
}

func (role *Role)TableName()string{
	return "roles"
}
