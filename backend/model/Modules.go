package model

import (
	"github.com/PoCFrance/CodeBaseManager/backend/database"
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

// ListModules: Return list of all Module from database.Database with associated Project ID..
func ListModules(project *Project) (modules []Module, err error) {
	if err = database.BackendDB.DB.Model(project).Related(&modules).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return modules, nil
}

// Find: Search for given Module.Name in database.Database with associated Project ID.
func (m *Module) FindByName(project *Project) (modules []Module, err error) {
	if err = database.BackendDB.DB.Model(project).Related(&modules).Where("name = ? ", m.Name).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return modules, nil
}

// Find: Search for given Module ID in database.Database with associated Project ID.
func (m *Module) FindById() (modules []Module, err error) {
	if err = database.BackendDB.DB.Model(m).Where("id = ?", m.ID).First(m).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return modules, nil
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
	if err := database.BackendDB.DB.Save(m).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return m, nil
}

// Delete: remove Module from database.Database
func (m *Module) Delete() (*Module, error) {
	if err := database.BackendDB.DB.Delete(m).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return m, nil
}
