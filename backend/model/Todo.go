package model

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Path string `gorm:"size:255;not null" json:"path"`
}

func (t *Todo) Save(db *database.Database) (*Todo, error) {
	if err := db.DB.Create(&t).Error; err != nil {
		return &Todo{}, err
	}
	return t, nil
}

func (t *Todo) Update(db *database.Database) (*Todo, error) {
	if err := db.DB.Update(&t).Error; err != nil {
		return &Todo{}, err
	}
	return t, nil
}

func (t *Todo) Delete(db *database.Database) (*Todo, error) {
	if err := db.DB.Delete(&t).Error; err != nil {
		return &Todo{}, err
	}
	return t, nil
}
