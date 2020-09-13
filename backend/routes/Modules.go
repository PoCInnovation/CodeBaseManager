package routes

import (
	"github.com/PoCFrance/CodeBaseManager/backend/controllers"
	"github.com/PoCFrance/CodeBaseManager/backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func addModule(c *gin.Context) {
	queryProject := &model.Project{}

	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryProject.ID = uint(projectId)

	queryModule := &model.Module{
		Name: c.Query("moduleName"),
		Path: c.Query("modulePath"),
	}

	if module, err := controllers.AddModule(queryProject, queryModule); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusCreated, module)
	}
}

func listModules(c *gin.Context) {
	queryProject := &model.Project{}

	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryProject.ID = uint(projectId)

	if modules, err := controllers.ListModules(queryProject); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, modules)
	}
}

func findModuleById(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryModule.ID = uint(moduleId)

	if module, err := controllers.FindModuleById(queryModule); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, module)
	}
}

func findModuleByName(c *gin.Context) {
	queryProject := &model.Project{}

	projectId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryProject.ID = uint(projectId)

	queryModule := &model.Module{
		Name: c.Param(rModule),
	}

	if modules, err := controllers.FindModuleByName(queryProject, queryModule); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, modules)
	}
}

func updateModule(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryModule.ID = uint(moduleId)

	updatedFields := &model.Module{
		Name: c.Query("moduleName"),
		Path: c.Query("modulePath"),
	}

	if module, err := controllers.UpdateModule(queryModule, updatedFields); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, module)
	}
}

func deleteModule(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryModule.ID = uint(moduleId)
	if module, err := controllers.DeleteModule(queryModule); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, module)
	}
}
