package model

import (
	"cbm-api/database"
	"github.com/jinzhu/gorm"
	"log"
)

// Module: Typedef for Module model.
type Module struct {
	gorm.Model
	Name      string `gorm:"size:255;not null" json:"name"`
	Path      string `gorm:"size:255;not null" json:"path"`
	ProjectID uint
	Functions []Function `json:"functions"`
	Types     []Type     `json:"types"`
}

// ListProjects: Return list of all Module from database.Database.
func ListModules(project *Project) (modules []Module, err error) {
	if err = database.BackendDB.DB.Model(project).Related(&modules).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return modules, nil
}

// Find: Search for given Project in database.Database with associated Project
func (m *Module) Find(project *Project) (*Module, error) {
	var modules []Module
	if err := database.BackendDB.DB.Model(project).Related(&modules).Where("name = ? ", m.Name).First(m).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return m, nil
}

// Save: create Module into database.Database with associated Project
func (m *Module) Save(project *Project) (*Module, error) {
	if err := database.BackendDB.DB.Model(project).Association("Modules").Append(m).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return m, nil
}

// Update: update Module from database.Database
func (m *Module) Update() (*Module, error) {
	if err := database.BackendDB.DB.Update(m).Error; err != nil {
		log.Print(err)
		return &Module{}, err
	}
	return m, nil
}

// Delete: remove Module from database.Database
func (m *Module) Delete() (*Module, error) {
	if err := database.BackendDB.DB.Delete(m).Error; err != nil {
		log.Print(err)
		return &Module{}, err
	}
	return m, nil
}
