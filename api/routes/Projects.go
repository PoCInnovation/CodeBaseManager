package routes

import (
	"cbm-api/database"
	"cbm-api/models_v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func addProject(c *gin.Context) {
	log.Print(c.Request)
	c.String(http.StatusOK, "list of all modules in "+c.Param(rProject))
}

func listProject(c *gin.Context) {
	result := database.CbmDb.DB.Find(&models_v2.Project{})
	if result.Error != nil {
		c.Value(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}
