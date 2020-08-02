package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
)

func FindProject(db *database.Database, name string) (*models.Project, error) {
	project := models.Project{
		Name: name,
	}
	return project.Find(db)
}

func ListProjects(db *database.Database) ([]models.Project, error) {
	return models.ListProject(db)
}

func DeleteProject(db *database.Database, name string) (project *models.Project, err error) {
	if project, err = FindProject(db, name); err != nil {
		return nil, err
	}
	if project, err = project.Delete(db); err != nil {
		return nil, err
	}
	return project, nil
}
