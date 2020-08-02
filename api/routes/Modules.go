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
	newModule := models_v2.Module{
		Name: name,
		Path: path,
	}
	var err error
	if _, err = newModule.SaveModule(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newModule)
	}
}

func listModule(c *gin.Context) {
	result := database.CbmDb.DB.Find(&models_v2.Module{})
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func findModule(c *gin.Context) {
	name := c.Param(rModule)
	module := models_v2.Module{
		Name: name,
	}

	result := database.CbmDb.DB.First(&module)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func deleteModule(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")

	module := models_v2.Module{
		Name: name,
		Path: path,
	}
	//var err error

	result := database.CbmDb.DB.First(&module)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
		return
	}
	if _, err := module.DeleteModule(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}