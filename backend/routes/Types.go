package routes

import (
	"cbm-api/database"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addType(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")
	newType := models.Type{
		Name: name,
		Path: path,
	}
	var err error
	if _, err = newType.SaveType(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	} else {
		c.JSON(http.StatusCreated, newType)
	}
}

func listType(c *gin.Context) {
	result := database.CbmDb.DB.Find(&models.Type{})
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func findType(c *gin.Context) {
	name := c.Param(rType)
	toFindType := models.Type{
		Name: name,
	}

	result := database.CbmDb.DB.First(&toFindType)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func deleteType(c *gin.Context) {
	name := c.Query("name")
	path := c.Query("path")

	toFindType := models.Type{
		Name: name,
		Path: path,
	}
	//var err error

	result := database.CbmDb.DB.First(&toFindType)
	if result.Error != nil {
		c.Value(http.StatusNotFound)
		return
	}
	if _, err := toFindType.DeleteType(database.CbmDb.DB); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
