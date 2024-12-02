package database

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate()
	log.Println("Database migrated successfully")
}
