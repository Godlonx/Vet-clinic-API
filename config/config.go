package config

import (
	"clinic/database"
	"clinic/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	CatRepository       dbmodel.CatRepository
	VisitRepository     dbmodel.VisitRepository
	TreatmentRepository dbmodel.TreatmentRepository
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)

	config.CatRepository = dbmodel.NewCatRepository(databaseSession)
	config.VisitRepository = dbmodel.NewVisitRepository(databaseSession)
	config.TreatmentRepository = dbmodel.NewTreatmentRepository(databaseSession)

	return &config, nil
}
