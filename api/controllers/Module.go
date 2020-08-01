package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
	"github.com/jinzhu/gorm"
)

func FindModule(db database.Database, project *models.Project, name string) *models.Module {
	module := models.Module{
		Name: name,
	}
	result := db.DB.Model(project).Related(&module)
	if result.Error != nil {
		return nil
	}
	return &module
}

func ListModule(db database.Database, project *models.Project) *gorm.DB {
	result := db.DB.Model(project).Find(&models.Module{})
	if result.Error != nil {
		return nil
	}
	return result
}
