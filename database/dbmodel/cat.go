package dbmodel

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	// ID    int    `json:"id" gorm:"primary_key"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
	Name  string `json:"name"`
	Weight float64 `json:"weight"`
}

type CatRepository interface {
	Create(cat *Cat) (*Cat, error)
	FindAll() ([]*Cat, error)
	Find(id int) (*Cat, error)
	Update(cat *Cat) (*Cat, error)
	Delete(id int) error
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Create(cat *Cat) (*Cat, error) {
	if err := r.db.Create(cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *catRepository) FindAll() ([]*Cat, error) {
	var cats []*Cat
	if err := r.db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *catRepository) Find(id int) (*Cat, error) {
	var cat Cat
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *catRepository) Update(cat *Cat) (*Cat, error) {
	if err := r.db.Save(cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *catRepository) Delete(id int) error {
	if err := r.db.Delete(&Cat{}, id).Error; err != nil {
		return err
	}
	return nil
}
