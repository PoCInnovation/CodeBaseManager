package controllers

import (
	"cbm-api/model"
	"errors"
	"fmt"
)

// AddModule: add a model.Project with given projectId, moduleName and modulePath.
//  Return an error if error in database.Database save.
func AddModule(project *model.Project, module *model.Module) (*model.Module, error) {
	if _, err := project.FindById(); err != nil {
		return nil, errors.New("project " + project.Name + " not found")
	}
	return module.Save(project)
}

// FindProjectByName : search for model.Project with FindByName method.
//  Return an error if no project found or error in database.Database query.
func FindModuleByName(project *model.Project, module *model.Module) (modules []model.Module, err error) {
	if _, err = project.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("project %d not found", project.ID))
	}
	if modules, err = module.FindByName(project); err != nil {
		return nil, errors.New("module " + module.Name + " not found")
	}
	if len(modules) == 0 {
		return nil, errors.New("no module found")
	}
	return modules, nil
}

// FindModuleById : search for model.Module with model.Module FindById method.
//  Return an error if no project or module found, or error in database.Database query.
func FindModuleById(module *model.Module) (*model.Module, error) {
	if _, err := module.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("module %d not found", module.ID))
	}
	return module, nil
}

//func FindModuleById(project *model.Project, module *model.Module) (*model.Module, error) {
//	if _, err := project.FindById(); err != nil {
//		return nil, errors.New("project " + project.Name + " not found")
//	}
//	if _, err := module.Find(project); err != nil {
//		return nil, errors.New("module " + module.Name + " not found")
//	}
//	return module, nil
//}

func ListModules(project *model.Project) ([]model.Module, error) {
	if _, err := project.FindByName(); err != nil {
		return nil, errors.New("project " + project.Name + " not found")
	}
	modules, err := model.ListModules(project)
	if err != nil {
		return nil, err
	}
	if len(modules) == 0 {
		return nil, errors.New("no modules found")
	}
	return modules, nil
}

func DeleteModule(module *model.Module) (*model.Module, error) {
	var err error
	if module, err = FindModuleById(module); err != nil {
		return nil, err
	}
	if module, err = module.Delete(); err != nil {
		return nil, err
	}
	return module, nil
}
