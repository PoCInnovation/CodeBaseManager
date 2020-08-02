package routes

import (
	"cbm-api/database"
	"cbm-api/models_v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addFunction(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	newFunction := models_v2.Function{
		Name: name,
		Path: path,
	}
	var err error
	if _, err = newFunction.SaveFunction(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newFunction)
	}
}

func listFunction(c *gin.Context) {
	result := database.CbmDb.DB.Find(&models_v2.Function{})
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func findFunction(c *gin.Context) {
	name := c.Param(rFunction)
	function := models_v2.Function{
		Name: name,
	}

	result := database.CbmDb.DB.First(&function)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func deleteFunction(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")

	function := models_v2.Function{
		Name: name,
		Path: path,
	}
	//var err error

	result := database.CbmDb.DB.First(&function)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
		return
	}
	if _, err := function.DeleteFunction(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}