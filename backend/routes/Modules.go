package routes

import (
	"github.com/PoCInnovation/CodeBaseManager/backend/controllers"
	"github.com/PoCInnovation/CodeBaseManager/backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// addModule: Call controllers.AddModule with projectId, moduleName and modulePath Query params.
// Abort when error occurs.
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

// listModules: Call controllers.ListModules with projectId Query param.
// Abort when error occurs.
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

// findModuleById: Call controllers.FindModuleById with Query param.
// Abort when error occurs.
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

// findModuleByName: Call controllers.FindModuleByName with module name in URL param and "projectId" in Query param.
// Abort when error occurs.
func findModuleByName(c *gin.Context) {
	queryProject := &model.Project{}

	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 64)
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

// findModuleByPath: Call controllers.FindModuleByPath with modulePath Query param.
// Abort when error occurs.
func findModuleByPath(c *gin.Context) {
	queryModule := &model.Module{
		Path: c.Query("modulePath"),
	}

	if modules, err := controllers.FindModuleByPath(queryModule); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, modules)
	}
}

// updateModule: Call controllers.UpdateProject with moduleId Query param.
// Abort when error occurs.
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

// deleteModule: Call controllers.DeleteModule with moduleId Query param.
// Abort when error occurs.
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
