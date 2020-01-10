package DB

import (
	"gorm"
)


/*
	开启 mysql 事务操作
	支持一次传递多个 *gorm.DB 执行语句(exce)
 */
func Transaction(dbs ...*gorm.DB) bool {

	tx := DB.Begin()

	for _,db := range dbs{
		tx = db
		if tx.Error != nil {
			tx.Rollback()
			return false
		}
	}

	tx.Commit()
	return true
}