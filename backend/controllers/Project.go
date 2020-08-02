package controllers

import (
	"cbm-api/models"
	"errors"
)

func FindProject(project *models.Project) (*models.Project, error) {
	if _, err := project.Find(); err != nil {
		return nil, errors.New("project " + project.Name + " not found")
	}
	return project, nil
}

func ListProjects() ([]models.Project, error) {
	projects, err := models.ListProject()
	if err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, errors.New("no project found")
	}
	return projects, nil
}

func AddProject(project *models.Project) (*models.Project, error) {
	return project.Save()
}

func DeleteProject(project *models.Project) (*models.Project, error) {
	var err error
	if project, err = FindProject(project); err != nil {
		return nil, err
	}
	if project, err = project.Delete(); err != nil {
		return nil, err
	}
	return project, nil
}
