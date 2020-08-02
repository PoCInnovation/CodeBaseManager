package models_v2

import "github.com/jinzhu/gorm"

type Type struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Path string `gorm:"size:255;not null;unique" json:"path"`
}

func (t *Type) SaveType(db *gorm.DB) (*Type, error) {
	if err := db.Create(&t).Error; err != nil {
		return &Type{}, err
	}
	return t, nil
}

func (t *Type) UpdateType(db *gorm.DB) (*Type, error) {
	if err := db.Update(&t).Error; err != nil {
		return &Type{}, err
	}
	return t, nil
}

func (t *Type) DeleteType(db *gorm.DB) (*Type, error) {
	if err := db.Delete(&t).Error; err != nil {
		return &Type{}, err
	}
	return t, nil
}