package model

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
	"log"
)

// Project: Typedef for Project model in database.Database
type Project struct {
	gorm.Model
	Name    string   `gorm:"size:255;not null" json:"name"` // Name: name of the Project
	Path    string   `gorm:"size:255;not null" json:"path"` // Path: Path of the Project
	Modules []Module `json:"modules"`                       // Modules: List of Project's related Module
	Todos   []Todo   `json:"todo"`                          // Todos: List of Project's related Todo
}

// ListProject: Return list of all Project from database.Database
func ListProject() (projects []Project, err error) {
	if err = database.BackendDB.DB.Find(&projects).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return projects, nil
}

// Find: Search for given Project in database.Database
func (p *Project) Find() (*Project, error) {
	if err := database.BackendDB.DB.Where("name = ?", p.Name).First(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}

// Save: create Project into database.Database
func (p *Project) Save() (*Project, error) {
	if err := database.BackendDB.DB.Create(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}

// Update: update Project from database.Database
func (p *Project) Update() (*Project, error) {
	if err := database.BackendDB.DB.Update(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}

// Delete: remove Project from database.Database
func (p *Project) Delete() (*Project, error) {
	if err := database.BackendDB.DB.Delete(p).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}
