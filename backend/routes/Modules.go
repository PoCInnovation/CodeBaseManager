package routes

import (
	"cbm-api/controllers"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addModule(c *gin.Context) {
	queryProject := &models.Project{
		Name: c.Query("projectName"),
	}
	queryModule := &models.Module{
		Name: c.Query("moduleName"),
		Path: c.Query("modulePath"),
	}

	if module, err := controllers.AddModule(queryProject, queryModule); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, module)
	}
}

func listModules(c *gin.Context) {
	queryProject := &models.Project{
		Name: c.Query("projectName"),
	}
	if modules, err := controllers.ListModules(queryProject); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, modules)
	}
}

func findModule(c *gin.Context) {
	queryProject := &models.Project{
		Name: c.Query("projectName"),
	}
	queryModule := &models.Module{
		Name: c.Query("moduleName"),
	}

	if module, err := controllers.FindModule(queryProject, queryModule); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusOK, module)
	}
}

func deleteModule(c *gin.Context) {
	queryProject := &models.Project{
		Name: c.Query("projectName"),
	}
	queryModule := &models.Module{
		Name: c.Query("moduleName"),
	}
	if module, err := controllers.DeleteModule(queryProject, queryModule); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusOK, module)
	}
}
