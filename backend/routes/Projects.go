package routes

import (
	"github.com/PoCInnovation/CodeBaseManager/backend/controllers"
	"github.com/PoCInnovation/CodeBaseManager/backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// addProject: Call controllers.AddProject with projectName and projectPath Query params.
// Abort when error occurs.
func addProject(c *gin.Context) {
	newProject := &model.Project{
		Name: c.Query("projectName"),
		Path: c.Query("projectPath"),
	}
	if project, err := controllers.AddProject(newProject); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusCreated, project)
	}
}

// listProject: Call controllers.ListProject.
// Abort when error occurs.
func listProject(c *gin.Context) {
	if projects, err := controllers.ListProjects(); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, projects)
	}
}

// findProjectByName: Call controllers.FindProjectByName with URL param.
// Abort when error occurs.
func findProjectByName(c *gin.Context) {
	queryProject := &model.Project{
		Name: c.Param(RProject),
	}
	if project, err := controllers.FindProjectByName(queryProject); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

// findProjectByPath: Call controllers.FindProjectByPath with projectPath Query param.
// Abort when error occurs.
func findProjectByPath(c *gin.Context) {
	queryProject := &model.Project{
		Path: c.Query("projectPath"),
	}
	if project, err := controllers.FindProjectByPath(queryProject); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

// findProjectById: Call controllers.FindProjectById with projectId Query param.
// Abort when error occurs.
func findProjectById(c *gin.Context) {
	queryProject := &model.Project{}

	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryProject.ID = uint(projectId)

	if project, err := controllers.FindProjectById(queryProject); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

// updateProject: Call controllers.UpdateProject with projectId Query param.
// Abort when error occurs.
func updateProject(c *gin.Context) {
	queryProject := &model.Project{}

	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryProject.ID = uint(projectId)

	updatedFields := &model.Project{
		Name: c.Query("projectName"),
		Path: c.Query("projectPath"),
	}

	if project, err := controllers.UpdateProject(queryProject, updatedFields); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

// deleteProject: Call controllers.DeleteProject with projectId Query param.
// Abort when error occurs.
func deleteProject(c *gin.Context) {
	queryProject := &model.Project{}

	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryProject.ID = uint(projectId)

	if project, err := controllers.DeleteProject(queryProject); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, project)
	}
}
