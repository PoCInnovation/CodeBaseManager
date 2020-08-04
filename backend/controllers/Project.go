package controllers

import (
	"cbm-api/model"
	"errors"
	"fmt"
)

// FindProjectByName : search for model.Project with FindByName method.
//  If no project found, return an error.
func FindProjectByName(project *model.Project) (projects []model.Project, err error) {
	if projects, err = project.FindByName(); err != nil {
		return nil, errors.New("project " + project.Name + " not found.")
	}
	if len(projects) == 0 {
		return nil, errors.New("no project found")
	}
	return projects, nil
}

// FindProjectById : search for model.Project with model.Project FindById method.
//  If no project found, return an error.
func FindProjectById(project *model.Project) (*model.Project, error) {
	if _, err := project.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("project %d not found", project.ID))
	}
	return project, nil
}

// ListProjects: search for list of all model.Project with model.ListProjects.
//  If no project found, return an error.
func ListProjects() ([]model.Project, error) {
	projects, err := model.ListProjects()
	if err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, errors.New("no project found")
	}
	return projects, nil
}

// AddProject: add a model.Project with given name and path.
func AddProject(project *model.Project) (*model.Project, error) {
	return project.Save()
}

// DeleteProject: Delete a model.Project with given Id.
func DeleteProject(project *model.Project) (*model.Project, error) {
	var err error
	if project, err = FindProjectById(project); err != nil {
		return nil, err
	}
	if project, err = project.Delete(); err != nil {
		return nil, err
	}
	return project, nil
}
