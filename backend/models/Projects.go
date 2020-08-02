package models

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Path string `gorm:"size:255;not null" json:"path"`
	//Name    string   `gorm:"size:255;not null;unique" json:"name"`
	//Path    string   `gorm:"size:255;not null;unique" json:"path"`
	Modules []Module `json:"modules"`
	Todos   []Module `json:"todo"`
}

func ListProject(db *database.Database) (projects []Project, err error) {
	if err = db.DB.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (p *Project) Save(db *database.Database) (*Project, error) {
	if err := db.DB.Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Project) Find(db *database.Database) (*Project, error) {
	if err := db.DB.First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Project) Update(db *database.Database) (*Project, error) {
	if err := db.DB.Update(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Project) Delete(db *database.Database) (*Project, error) {
	if err := db.DB.Delete(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
