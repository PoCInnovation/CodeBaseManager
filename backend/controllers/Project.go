package controllers

import (
	"cbm-api/model"
	"errors"
)

// FindProject : Find with ProjectFind
func FindProject(project *model.Project) (*model.Project, error) {
	if _, err := project.Find(); err != nil {
		return nil, errors.New("project " + project.Name + " not found")
	}
	return project, nil
}

func ListProjects() ([]model.Project, error) {
	projects, err := model.ListProject()
	if err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, errors.New("no project found")
	}
	return projects, nil
}

func AddProject(project *model.Project) (*model.Project, error) {
	return project.Save()
}

func DeleteProject(project *model.Project) (*model.Project, error) {
	var err error
	if project, err = FindProject(project); err != nil {
		return nil, err
	}
	if project, err = project.Delete(); err != nil {
		return nil, err
	}
	return project, nil
}
