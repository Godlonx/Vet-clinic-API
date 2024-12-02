package database

import (
	"clinic/database/dbmodel"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.Cat{},
	)
	log.Println("Database migrated successfully")
}
