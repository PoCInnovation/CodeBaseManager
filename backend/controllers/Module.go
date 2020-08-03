package controllers

import (
	"cbm-api/models"
	"errors"
)

func AddModule(project *models.Project, module *models.Module) (*models.Module, error) {
	if _, err := project.Find(); err != nil {
		return nil, errors.New("project " + project.Name + " not found")
	}
	return module.Save(project)
}

func FindModule(project *models.Project, module *models.Module) (*models.Module, error) {
	if _, err := project.Find(); err != nil {
		return nil, errors.New("project " + project.Name + " not found")
	}
	if _, err := module.Find(project); err != nil {
		return nil, errors.New("module " + module.Name + " not found")
	}
	return module, nil
}

func ListModules(project *models.Project) ([]models.Module, error) {
	if _, err := project.Find(); err != nil {
		return nil, errors.New("project " + project.Name + " not found")
	}
	modules, err := models.ListModules(project)
	if err != nil {
		return nil, err
	}
	if len(modules) == 0 {
		return nil, errors.New("no modules found")
	}
	return modules, nil
}

func DeleteModule(project *models.Project, module *models.Module) (*models.Project, error) {
	var err error
	if module, err = FindModule(project, module); err != nil {
		return nil, err
	}
	if module, err = module.Delete(); err != nil {
		return nil, err
	}
	return project, nil
}
