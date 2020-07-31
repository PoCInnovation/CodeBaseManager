package models_v2

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Path string `gorm:"size:255;not null" json:"path"`
	//Name    string   `gorm:"size:255;not null;unique" json:"name"`
	//Path    string   `gorm:"size:255;not null;unique" json:"path"`
	Modules []Module `json:"modules"`
}

func (p *Project) SaveProject(db *gorm.DB) (*Project, error) {
	if err := db.Create(&p).Error; err != nil {
		return &Project{}, err
	}
	return p, nil
}

func (p *Project) UpdateProject(db *gorm.DB) (*Project, error) {
	if err := db.Update(&p).Error; err != nil {
		return &Project{}, err
	}
	return p, nil
}

func (p *Project) DeleteProject(db *gorm.DB) (*Project, error) {
	if err := db.Delete(&p).Error; err != nil {
		return &Project{}, err
	}
	return p, nil
}
