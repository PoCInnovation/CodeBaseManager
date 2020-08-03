package model

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
)

type Function struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Path string `gorm:"size:255;not null;unique" json:"path"`
}

func (f *Function) Save(db *database.Database) (*Function, error) {
	if err := db.DB.Create(&f).Error; err != nil {
		return &Function{}, err
	}
	return f, nil
}

func (f *Function) Update(db *database.Database) (*Function, error) {
	if err := db.DB.Update(&f).Error; err != nil {
		return &Function{}, err
	}
	return f, nil
}

func (f *Function) Delete(db *database.Database) (*Function, error) {
	if err := db.DB.Delete(&f).Error; err != nil {
		return &Function{}, err
	}
	return f, nil
}
