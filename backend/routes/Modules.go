package routes

import (
	"cbm-api/controllers"
	"cbm-api/database"
	"cbm-api/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addModule(c *gin.Context) {
	projectName := c.Param(rProject)
	project := controllers.FindProject(database.CbmDb, projectName)
	if project == nil {
		c.String(http.StatusNotFound, "project not found")
		return
	}

	name := c.Query("name")
	path := c.Query("path")
	newModule := models.Module{
		Name: name,
		Path: path,
	}

	if _, err := newModule.Save(database.CbmDb.DB); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newModule)
	}
}

func listModules(c *gin.Context) {
	projectName := c.Param(rProject)
	project := controllers.FindProject(database.CbmDb, projectName)
	if project == nil {
		_ = c.AbortWithError(http.StatusNotFound, errors.New("project "+projectName+" not found"))
	}
	modules := controllers.ListModule(database.CbmDb, project)
	if modules == nil {
		_ = c.AbortWithError(http.StatusNotFound, errors.New("no modules found for project "+projectName))
	}
	c.JSON(http.StatusOK, modules)
}

func findModule(c *gin.Context) {
	projectName := c.Param(rProject)
	project := controllers.FindProject(database.CbmDb, projectName)
	if project == nil {
		_ = c.AbortWithError(http.StatusNotFound, errors.New("project "+projectName+" not found"))
	}
	module := controllers.FindModule(database.CbmDb, project, "TODO")
	c.JSON(http.StatusOK, module)
}

func deleteModule(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")

	project := models.Project{
		Name: name,
		Path: path,
	}
	//var err error

	result := database.CbmDb.DB.First(&project)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
		return
	}
	if _, err := project.Delete(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
