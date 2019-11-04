package model

/*
CREATE TABLE `tags` (
  `t_id` int(11) NOT NULL AUTO_INCREMENT,
  `t_name` char(20) DEFAULT NULL,
  PRIMARY KEY (`t_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
 */
// 标签
type Tag struct {
	TID  int    `gorm:"column:t_id" json:"tid"`
	Name string `gorm:"column:t_name" json:"tag"`
}

func (tg *Tag)TableName()string{
	return "tags"
}