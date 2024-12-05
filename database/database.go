package database

import (
	"clinic/database/dbmodel"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.Cat{},
		&dbmodel.Visit{},
		&dbmodel.Treatment{},
	)
	log.Println("Database migrated successfully")
}
