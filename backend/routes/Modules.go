package routes

import (
	"github.com/gin-gonic/gin"
)

func addModule(c *gin.Context) {
	//projectName := c.Param(rProject)
	//db := c.MustGet("db").(*database.Database)
	//project, err := controllers.FindProject(db, projectName)
	//if err != nil {
	//	_ = c.AbortWithError(http.StatusNotFound, errors.New("project "+projectName+" not found"))
	//}
	//
	//name := c.Query("name")
	//path := c.Query("path")
	//newModule := models.Module{
	//	Name: name,
	//	Path: path,
	//}
	//
	//if _, err := newModule.Save(database.CbmDb.DB); err != nil {
	//	_ = c.AbortWithError(http.StatusForbidden, err)
	//} else {
	//	c.JSON(http.StatusCreated, newModule)
	//}
}

func listModules(c *gin.Context) {
	//projectName := c.Param(rProject)
	//project := controllers.FindProject(database.CbmDb, projectName)
	//if project == nil {
	//	_ = c.AbortWithError(http.StatusNotFound, errors.New("project "+projectName+" not found"))
	//}
	//modules := controllers.ListModule(database.CbmDb, project)
	//if modules == nil {
	//	_ = c.AbortWithError(http.StatusNotFound, errors.New("no modules found for project "+projectName))
	//}
	//c.JSON(http.StatusOK, modules)
}

func findModule(c *gin.Context) {
	//projectName := c.Param(rProject)
	//project := controllers.FindProject(database.CbmDb, projectName)
	//if project == nil {
	//	_ = c.AbortWithError(http.StatusNotFound, errors.New("project "+projectName+" not found"))
	//}
	//module := controllers.FindModule(database.CbmDb, project, "TODO")
	//c.JSON(http.StatusOK, module)
}

func deleteModule(c *gin.Context) {
	//name := c.Query("name")
	//path := c.Query("path")
	//db := c.MustGet("db").(*database.Database)
	//
	//project := models.Project{
	//	Name: name,
	//	Path: path,
	//}
	////var err error
	//
	//result := database.CbmDb.DB.First(&project)
	//if result.Error != nil {
	//	c.Value(http.StatusNotFound)
	//	return
	//}
	//if _, err := project.Delete(db); err != nil {
	//	_ = c.AbortWithError(http.StatusNotFound, err)
	//} else {
	//	c.JSON(http.StatusOK, result)
	//}
}
