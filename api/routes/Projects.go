package routes

import (
	"cbm-api/controllers"
	"cbm-api/database"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addProject(c *gin.Context) {
	//TODO: move directly into structure ?
	name := c.Query("name")
	path := c.Query("path")
	newProject := models.Project{
		Name: name,
		Path: path,
	}

	if _, err := newProject.Save(database.CbmDb.DB); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newProject)
	}
}

func listProject(c *gin.Context) {
	projects := controllers.ListProjects(database.CbmDb)

	if projects == nil {
		c.Value(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, projects)
	//var projects []models.Project
	//result := database.CbmDb.DB.Find(&projects)
	//if result.Error != nil {
	//	c.Value(http.StatusNotFound)
	//}
	//c.JSON(http.StatusOK, projects)
}

func findProject(c *gin.Context) {
	project := controllers.FindProject(database.CbmDb, c.Param(rProject))

	if project == nil {
		c.Value(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, project)
}

func deleteProject(c *gin.Context) {
	//TODO: move directly into structure ?
	name := c.Query("name")
	path := c.Query("path")
	project := models.Project{
		Name: name,
		Path: path,
	}

	result := database.CbmDb.DB.First(&project)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
		return
	}
	if _, err := project.Delete(database.CbmDb.DB); err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
