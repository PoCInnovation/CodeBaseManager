package routes

import (
	"cbm-api/controllers"
	"cbm-api/database"
	"cbm-api/models"
	"errors"
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
	db := c.MustGet("db").(*database.Database)

	if _, err := newProject.Save(db); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newProject)
	}
}

func listProject(c *gin.Context) {
	db := c.MustGet("db").(*database.Database)
	projects, err := controllers.ListProjects(db)

	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, errors.New("no project found"))
	}
	c.JSON(http.StatusOK, projects)
}

func findProject(c *gin.Context) {
	projectName := c.Query("name")
	db := c.MustGet("db").(*database.Database)
	project, err := controllers.FindProject(db, projectName)

	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, errors.New("project "+projectName+" not found"))
	}
	c.JSON(http.StatusOK, project)
}

func deleteProject(c *gin.Context) {
	projectName := c.Query("name")
	//TODO: move directly into structure ?
	db := c.MustGet("db").(*database.Database)
	if project, err := controllers.DeleteProject(db, projectName); err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, project)
	}

}
