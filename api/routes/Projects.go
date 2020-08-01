package routes

import (
	"cbm-api/controllers"
	"cbm-api/database"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addProject(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	newProject := models.Project{
		Name: name,
		Path: path,
	}

	if _, err := newProject.Save(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newProject)
	}
}

func listProject(c *gin.Context) {
	var projects []models.Project
	result := database.CbmDb.DB.Find(&projects)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, projects)
	//c.String(http.StatusOK, "List of all project")
}

func findProject(c *gin.Context) {
	projectName := c.Param(rProject)
	project := controllers.FindProject(database.CbmDb, projectName)

	if project == nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, project)
	//project := models.Project{
	//	Name: name,
	//}
	//
	//result := database.CbmDb.DB.First(&project)
	//if result.Error != nil {
	//	c.Value(http.StatusNotFound)
	//}
	//c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func deleteProject(c *gin.Context) {
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
