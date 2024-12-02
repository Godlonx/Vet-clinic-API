package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	// ID     int       `json:"id" gorm:"primary_key"`
	CatId  int       `json:"cat_id"`
	Date   time.Time `json:"date"`
	Reason string    `json:"reason"`
}

type VisitRepository interface {
	Create(visit Visit) error
	Find(id uint) (Visit, error)
	FindAll() ([]Visit, error)
	Update(visit Visit) error
	Delete(id uint) error
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db: db}
}

func (r *visitRepository) Create(visit Visit) error {
	if err := r.db.Create(&visit).Error; err != nil {
		return err
	}
	return nil
}

func (r *visitRepository) Find(id uint) (Visit, error) {
	var visit Visit
	if err := r.db.First(&visit, id).Error; err != nil {
		return visit, err
	}
	return visit, nil
}

func (r *visitRepository) FindAll() ([]Visit, error) {
	var visits []Visit
	if err := r.db.Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *visitRepository) Update(visit Visit) error {
	if err := r.db.Save(&visit).Error; err != nil {
		return err
	}
	return nil
}

func (r *visitRepository) Delete(id uint) error {
	if err := r.db.Delete(&Visit{}, id).Error; err != nil {
		return err
	}
	return nil
}
