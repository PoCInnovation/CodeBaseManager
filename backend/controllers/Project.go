package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
	"errors"
)

func FindProject(db *database.Database, project *models.Project) (*models.Project, error) {
	return project.Find(db)
}

func ListProjects(db *database.Database) ([]models.Project, error) {
	projects, err := models.ListProject(db)
	if err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, errors.New("no project found")
	}
	return projects, nil
}

func AddProject(db *database.Database, project *models.Project) (*models.Project, error) {
	return project.Save(db)
}

func DeleteProject(db *database.Database, project *models.Project) (*models.Project, error) {
	var err error
	if project, err = FindProject(db, project); err != nil {
		return nil, err
	}
	if project, err = project.Delete(db); err != nil {
		return nil, err
	}
	return project, nil
}
