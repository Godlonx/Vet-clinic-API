package dbmodel

import (
	"gorm.io/gorm"
)

// Treatment represents a treatment given to a pet
type Treatment struct {
	gorm.Model
	// ID          int       `json:"id" gorm:"primary_key"`
	VisitID     int     `json:"visit_id"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Cost        float64 `json:"cost"`
}

// TreatmentRepository defines the interface for treatment operations
type TreatmentRepository interface {
	Create(treatment *Treatment) error
	FindAll() ([]*Treatment, error)
	Find(id int) (*Treatment, error)
	FindByVisit(id int) (*Treatment, error)
	Update(treatment *Treatment) (*Treatment, error)
	Delete(id int) error
}

type treatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRepository(db *gorm.DB) TreatmentRepository {
	return &treatmentRepository{db: db}
}

func (r *treatmentRepository) Create(treatment *Treatment) error {
	if err := r.db.Create(treatment).Error; err != nil {
		return err
	}
	return nil
}

func (r *treatmentRepository) FindAll() ([]*Treatment, error) {
	var treatments []*Treatment
	if err := r.db.Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}

func (r *treatmentRepository) Find(id int) (*Treatment, error) {
	var treatment Treatment
	if err := r.db.First(&treatment, id).Error; err != nil {
		return nil, err
	}
	return &treatment, nil
}

func (r *treatmentRepository) FindByVisit(id int) (*Treatment, error) {
	var treatment Treatment
	req := r.db.Where("visit_id = ?", id)
	if err := req.First(&treatment).Error; err != nil {
		return nil, err
	}
	return &treatment, nil
}

func (r *treatmentRepository) Update(treatment *Treatment) (*Treatment, error) {
	if err := r.db.Save(treatment).Error; err != nil {
		return nil, err
	}
	return treatment, nil
}

func (r *treatmentRepository) Delete(id int) error {
	if err := r.db.Delete(&Treatment{}, id).Error; err != nil {
		return err
	}
	return nil
}
