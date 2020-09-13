package model

import (
	"github.com/PoCFrance/CodeBaseManager/backend/database"
	"github.com/jinzhu/gorm"
	"log"
)

type Function struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	Path     string `gorm:"size:255;not null" json:"path"`
	ModuleID uint
}

// ListFunctions: Return list of all Function from database.Database with associated Module ID..
func ListFunctions(module *Module) (functions []Function, err error) {
	if err = database.BackendDB.DB.Model(module).Related(&functions).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return functions, nil
}

// Find: Search for given Function.Name in database.Database with associated Module ID.
func (f *Function) FindByName(module *Module) (functions []Function, err error) {
	if err = database.BackendDB.DB.Model(module).Related(&functions).Where("name = ? ", f.Name).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return functions, nil
}

// Find: Search for given Function ID in database.Database with associated Module ID.
func (f *Function) FindById() (functions []Function, err error) {
	if err = database.BackendDB.DB.Model(f).Where("id = ?", f.ID).First(f).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return functions, nil
}

// Save: create Function into database.Database with associated Module
func (f *Function) Save(module *Module) (*Function, error) {
	if err := database.BackendDB.DB.Model(module).Association("Functions").Append(f).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return f, nil
}

// Update: update Function from database.Database
func (f *Function) Update() (*Function, error) {
	if err := database.BackendDB.DB.Save(f).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return f, nil
}

// Delete: remove Module from database.Database
func (f *Function) Delete() (*Function, error) {
	if err := database.BackendDB.DB.Delete(f).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return f, nil
}
