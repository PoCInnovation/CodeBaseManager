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

// ListProjects: Return list of all Module from database.Database with associated Project ID..
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

//// FindByName: Search for list of Project in database.Database with Project.Name
//func (p *Project) FindByName() (projects []Project, err error) {
//	if err = database.BackendDB.DB.Where("name = ?", p.Name).Find(&projects).Error; err != nil {
//		log.Print(err)
//		return nil, err
//	}
//	return projects, nil
//}

//// FindById: Search for list of Project in database.Database with Project ID
//func (p *Project) FindById() (*Project, error) {
//	if err := database.BackendDB.DB.Where("id = ?", p.ID).First(p).Error; err != nil {
//		log.Print(err)
//		return nil, err
//	}
//	return p, nil
//}

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
