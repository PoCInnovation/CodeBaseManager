package models_v2

import "github.com/jinzhu/gorm"

type Function struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Path string `gorm:"size:255;not null;unique" json:"path"`
}

func (f *Function) SaveFunction(db *gorm.DB) (*Function, error) {
	if err := db.Create(&f).Error; err != nil {
		return &Function{}, err
	}
	return f, nil
}

func (f *Function) UpdateFunction(db *gorm.DB) (*Function, error) {
	if err := db.Update(&f).Error; err != nil {
		return &Function{}, err
	}
	return f, nil
}

func (f *Function) DeleteFunction(db *gorm.DB) (*Function, error) {
	if err := db.Delete(&f).Error; err != nil {
		return &Function{}, err
	}
	return f, nil
}