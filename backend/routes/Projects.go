package routes

import (
	"github.com/PoCFrance/CodeBaseManager/backend/controllers"
	"github.com/PoCFrance/CodeBaseManager/backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func addProject(c *gin.Context) {
	log.Println(c.Request.Form)
	newProject := &model.Project{
		Name: c.Query("projectName"),
		Path: c.Query("projectPath"),
	}
	log.Println(newProject)
	if project, err := controllers.AddProject(newProject); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusCreated, project)
	}
}

func listProject(c *gin.Context) {
	if projects, err := controllers.ListProjects(); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, projects)
	}
}

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
