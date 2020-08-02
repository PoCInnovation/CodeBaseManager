package models

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
)

type Module struct {
	gorm.Model
	Name      string     `gorm:"size:255;not null;unique" json:"name"`
	Path      string     `gorm:"size:255;not null;unique" json:"path"`
	Functions []Function `json:"functions"`
	Types     []Type     `json:"types"`
}

func (m *Module) Save() (*Module, error) {
	if err := database.BackendDB.DB.Create(&m).Error; err != nil {
		return &Module{}, err
	}
	return m, nil
}

func (m *Module) Update() (*Module, error) {
	if err := database.BackendDB.DB.Create(&m).Error; err != nil {
		return &Module{}, err
	}
	return m, nil
}

func (m *Module) Delete() (*Module, error) {
	if err := database.BackendDB.DB.Delete(&m).Error; err != nil {
		return &Module{}, err
	}
	return m, nil
}
