package models

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
	"log"
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

func ListProject() (projects []Project, err error) {
	if err = database.BackendDB.DB.Find(&projects).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return projects, nil
}

func (p *Project) Save() (*Project, error) {
	if err := database.BackendDB.DB.Create(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}

func (p *Project) Find() (*Project, error) {
	if err := database.BackendDB.DB.Where("name = ?", p.Name).First(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}

func (p *Project) Update() (*Project, error) {
	if err := database.BackendDB.DB.Update(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}

func (p *Project) Delete() (*Project, error) {
	if err := database.BackendDB.DB.Delete(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}
