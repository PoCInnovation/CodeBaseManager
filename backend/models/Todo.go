package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Path string `gorm:"size:255;not null" json:"path"`
}

func (t *Todo) Save(db *gorm.DB) (*Todo, error) {
	if err := db.Create(&t).Error; err != nil {
		return &Todo{}, err
	}
	return t, nil
}

func (t *Todo) Update(db *gorm.DB) (*Todo, error) {
	if err := db.Update(&t).Error; err != nil {
		return &Todo{}, err
	}
	return t, nil
}

func (t *Todo) Delete(db *gorm.DB) (*Todo, error) {
	if err := db.Delete(&t).Error; err != nil {
		return &Todo{}, err
	}
	return t, nil
}
