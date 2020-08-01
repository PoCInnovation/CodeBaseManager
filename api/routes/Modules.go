package routes

import (
	"cbm-api/controllers"
	"cbm-api/database"
	"cbm-api/models"
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
		c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newModule)
	}
}

func listModule(c *gin.Context) {
	projectName := c.Param(rProject)
	project := controllers.FindProject(database.CbmDb, projectName)
	if project == nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, project)
}

func findModule(c *gin.Context) {
	name := c.Param(rProject)
	project := controllers.FindProject(database.CbmDb, name)

	if project == nil {
		c.String(http.StatusNotFound, "project not found")
	}

	result := database.CbmDb.DB.First(&project)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
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
