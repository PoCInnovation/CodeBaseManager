package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
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

func ListModule(db database.Database, project *models.Project) []models.Module {
	var modules []models.Module
	if err := db.DB.Model(project).Related(&modules).Error; err != nil {
		return nil
	}
	return modules
}
