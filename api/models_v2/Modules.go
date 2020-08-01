package models_v2

import "github.com/jinzhu/gorm"

type Module struct {
	gorm.Model
	Name      string     `gorm:"size:255;not null;unique" json:"name"`
	Path      string     `gorm:"size:255;not null;unique" json:"path"`
	Functions []Function `json:"functions"`
	Types     []Type     `json:"types"`
}

func (m *Module) Save(db *gorm.DB) (*Module, error) {
	if err := db.Create(&m).Error; err != nil {
		return &Module{}, err
	}
	return m, nil
}

func (m *Module) Update(db *gorm.DB) (*Module, error) {
	if err := db.Create(&m).Error; err != nil {
		return &Module{}, err
	}
	return m, nil
}

func (m *Module) Delete(db *gorm.DB) (*Module, error) {
	if err := db.Delete(&m).Error; err != nil {
		return &Module{}, err
	}
	return m, nil
}
