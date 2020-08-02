package routes

import (
	"cbm-api/controllers"
	"cbm-api/database"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addProject(c *gin.Context) {
	////TODO: move directly into structure ?
	//name := c.Query("name")
	//path := c.Query("path")
	db := c.MustGet("db").(*database.Database)
	newProject := &models.Project{
		Name: c.Query("name"),
		Path: c.Query("path"),
	}

	if project, err := controllers.AddProject(db, newProject); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, project)
	}
}

func listProject(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	projects, err := controllers.ListProjects(db)

	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, projects)
}

func findProject(c *gin.Context) {
	//projectName := c.Query("name")
	db := c.MustGet("db").(*database.Database)
	queryProject := &models.Project{
		Name: c.Query("name"),
	}

	if project, err := controllers.FindProject(db, queryProject); err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, project)
	}

}

func deleteProject(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	queryProject := &models.Project{
		Name: c.Query("name"),
		Path: c.Query("path"),
	}

	if project, err := controllers.DeleteProject(db, queryProject); err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, project)
	}

}
