package routes

import (
	"cbm-api/controllers"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addProject(c *gin.Context) {
	newProject := &models.Project{
		Name: c.Query("projectName"),
		Path: c.Query("projectPath"),
	}

	if project, err := controllers.AddProject(newProject); err != nil {
		_ = c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, project)
	}
}

func listProject(c *gin.Context) {
	if projects, err := controllers.ListProjects(); err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, projects)
	}
}

func findProject(c *gin.Context) {
	queryProject := &models.Project{
		Name: c.Query("projectName"),
	}
	if project, err := controllers.FindProject(queryProject); err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

func deleteProject(c *gin.Context) {
	queryProject := &models.Project{
		Name: c.Query("projectName"),
	}

	if project, err := controllers.DeleteProject(queryProject); err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}
