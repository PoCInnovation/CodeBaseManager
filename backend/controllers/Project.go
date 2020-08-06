package controllers

import (
	"cbm-api/model"
	"errors"
	"fmt"
)

// AddProject: add a model.Project with given projectName and projectPath.
//  Return an error if error in database.Database save.
func AddProject(project *model.Project) (*model.Project, error) {
	return project.Save()
}

// FindProjectByName : search for model.Project with FindByName method.
//  Return an error if no project found or error in database.Database query.
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
//  Return an error if no project found or error in database.Database query.
func FindProjectById(project *model.Project) (*model.Project, error) {
	if _, err := project.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("project %d not found", project.ID))
	}
	return project, nil
}

// UpdateProject : search for model.Project with model.Project FindById method and update fields.
//  Return an error if no project found or error in database.Database query or save.
func UpdateProject(project *model.Project) (*model.Project, error) {
	if _, err := project.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("project %d not found", project.ID))
	}
	return project, nil
}

// ListProjects: search for list of all model.Project with model.ListProjects.
//  Return an error if no project found or error in database.Database listing.
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

// DeleteProject: Delete a model.Project with given Id.
//  Return an error if no project found or error in database.Database deletion.
func DeleteProject(project *model.Project) (*model.Project, error) {
	var err error
	if project, err = FindProjectById(project); err != nil {
		return nil, err
	}
	if modules, err := ListModules(project); err == nil && modules != nil {
		for _, module := range modules {
			if _, err := module.Delete(); err != nil {
				return nil, err
			}
		}
	}
	if project, err = project.Delete(); err != nil {
		return nil, err
	}
	return project, nil
}
