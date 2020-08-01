package routes

import (
	"cbm-api/database"
	"cbm-api/models_v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addModule(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	newProject := models_v2.Project{
		Name: name,
		Path: path,
	}
	var err error
	if _, err = newProject.Save(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newProject)
	}
}

func listModule(c *gin.Context) {
	result := database.CbmDb.DB.Find(&models_v2.Project{})
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func findModule(c *gin.Context) {
	name := c.Param(rProject)
	project := models_v2.Project{
		Name: name,
	}

	result := database.CbmDb.DB.First(&project)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func deleteModule(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")

	project := models_v2.Project{
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
