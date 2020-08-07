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

// FindModuleByName : search for model.Module with FindByName method.
//  Return an error if no project or module found or error in database.Database query.
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

// UpdateModule : search for model.Module with model.Module FindById method and update fields.
//  Return an error if no Module found or error in database.Database query or save.
func UpdateModule(queryModule *model.Module, updatedFields *model.Module) (*model.Module, error) {
	if _, err := queryModule.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("Module %d not found", queryModule.ID))
	}
	if updatedFields.Path != "" {
		queryModule.Path = updatedFields.Path
	}
	if updatedFields.Name != "" {
		queryModule.Name = updatedFields.Name
	}
	if _, err := queryModule.Update(); err != nil {
		return nil, err
	}
	return queryModule, nil
}

// ListModules: search for list of all model.Model with associated model.ListProjects ID.
//  Return an error if no project found or error in database.Database listing.
func ListModules(project *model.Project) ([]model.Module, error) {
	if _, err := project.FindById(); err != nil {
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

func DeleteModuleDependencies(module *model.Module) (*model.Module, error) {
	if functions, err := ListFunctions(module); err == nil && functions != nil {
		for _, function := range functions {
			if _, err := function.Delete(); err != nil {
				return nil, err
			}
		}
	}
	return module, nil
}

func DeleteModule(module *model.Module) (*model.Module, error) {
	var err error
	if module, err = FindModuleById(module); err != nil {
		return nil, err
	}
	if module, err = DeleteModuleDependencies(module); err != nil {
		return nil, err
	}
	if module, err = module.Delete(); err != nil {
		return nil, err
	}
	return module, nil
}
