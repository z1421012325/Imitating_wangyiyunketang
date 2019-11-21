package DB

import (
	"gorm"
)



//  Transection 事务封装?
func Transaction(db *gorm.DB)bool{

	tx := DB.Begin()
	tx = db

	if tx.Error != nil {
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}