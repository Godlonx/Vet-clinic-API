package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

// Treatment represents a treatment given to a pet
type Treatment struct {
	gorm.Model
	// ID          int       `json:"id" gorm:"primary_key"`
	PetID       int       `json:"pet_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Cost        float64   `json:"cost"`
}

// TreatmentRepository defines the interface for treatment operations
type TreatmentRepository interface {
	Create(treatment *Treatment) error
	FindAll() ([]*Treatment, error)
	Find(id int) (*Treatment, error)
	Update(treatment *Treatment) (*Treatment, error)
	Delete(id int) error
}
