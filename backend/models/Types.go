package models

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
)

type Type struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Path string `gorm:"size:255;not null;unique" json:"path"`
}

func (t *Type) Save(db *database.Database) (*Type, error) {
	if err := db.DB.Create(&t).Error; err != nil {
		return &Type{}, err
	}
	return t, nil
}

func (t *Type) Update(db *database.Database) (*Type, error) {
	if err := db.DB.Update(&t).Error; err != nil {
		return &Type{}, err
	}
	return t, nil
}

func (t *Type) Delete(db *database.Database) (*Type, error) {
	if err := db.DB.Delete(&t).Error; err != nil {
		return &Type{}, err
	}
	return t, nil
}
