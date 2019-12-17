package DB

import (
	"gorm"
)



func Transaction(dbs ...*gorm.DB)bool{

	tx := DB.Begin()

	//tx = db
	//if tx.Error != nil {
	//	tx.Rollback()
	//	return false
	//}

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