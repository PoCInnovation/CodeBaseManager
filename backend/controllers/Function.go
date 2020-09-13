package controllers

import (
	"cbm-api/model"
	"errors"
	"fmt"
)

// AddFunction: add a model.Module with given moduleId, functionName and functionPath.
//  Return an error if error in database.Database save.
func AddFunction(module *model.Module, function *model.Function) (*model.Function, error) {
	if _, err := module.FindById(); err != nil {
		return nil, errors.New("module " + module.Name + " not found")
	}
	return function.Save(module)
}

// FindFunctionByName : search for model.Function with FindByName method.
//  Return an error if no module or function found or error in database.Database query.
func FindFunctionByName(module *model.Module, function *model.Function) (functions []model.Function, err error) {
	if _, err = module.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("module %d not found", module.ID))
	}
	if functions, err = function.FindByName(module); err != nil {
		return nil, errors.New("function " + function.Name + " not found")
	}
	if len(functions) == 0 {
		return nil, errors.New("no function found")
	}
	return functions, nil
}

// FindFunctionByPath : search for model.Function with FindByPath method.
//  Return an error if no module or function found or error in database.Database query.
func FindFunctionByPath(module *model.Module, function *model.Function) (functions []model.Function, err error) {
	if _, err = module.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("module %d not found", module.ID))
	}
	if functions, err = function.FindByPath(module); err != nil {
		return nil, fmt.Errorf("function with path {%s} not found", function.Path)
	}
	if len(functions) == 0 {
		return nil, errors.New("no function found")
	}
	return functions, nil
}

// FindFunctionById : search for model.Function with model.Function FindById method.
//  Return an error if no module or function found, or error in database.Database query.
func FindFunctionById(function *model.Function) (*model.Function, error) {
	if _, err := function.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("function %d not found", function.ID))
	}
	return function, nil
}

// UpdateFunction : search for model.Function with model.Function FindById method and update fields.
//  Return an error if no Function found or error in database.Database query or save.
func UpdateFunction(queryFunction *model.Function, updatedFields *model.Function) (*model.Function, error) {
	if _, err := queryFunction.FindById(); err != nil {
		return nil, errors.New(fmt.Sprintf("Function %d not found", queryFunction.ID))
	}
	if updatedFields.Path != "" {
		queryFunction.Path = updatedFields.Path
	}
	if updatedFields.Name != "" {
		queryFunction.Name = updatedFields.Name
	}
	if _, err := queryFunction.Update(); err != nil {
		return nil, err
	}
	return queryFunction, nil
}

// ListFunctions: search for list of all model.Model with associated model.ListModules ID.
//  Return an error if no module found or error in database.Database listing.
func ListFunctions(module *model.Module) ([]model.Function, error) {
	if _, err := module.FindById(); err != nil {
		return nil, errors.New("module " + module.Name + " not found")
	}
	functions, err := model.ListFunctions(module)
	if err != nil {
		return nil, err
	}
	if len(functions) == 0 {
		return nil, errors.New("no functions found")
	}
	return functions, nil
}

func DeleteFunction(function *model.Function) (*model.Function, error) {
	var err error
	if function, err = FindFunctionById(function); err != nil {
		return nil, err
	}
	if function, err = function.Delete(); err != nil {
		return nil, err
	}
	return function, nil
}
