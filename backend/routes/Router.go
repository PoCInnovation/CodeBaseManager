package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	rModule       = "module"
	rProject      = "project"
	rFunction     = "function"
	rType         = "type"
	rName         = "name"
	InternalError = http.StatusForbidden
)

//TODO:
// update project controller and model
// testing postman

func ApplyRoutes(r *gin.Engine) {

	projects := r.Group("/" + rProject)
	{
		projects.GET("/list", listProject)
		projects.GET("/get", findProjectById)
		projects.GET("/get/:"+rProject, findProjectByName)
		projects.POST("/add", addProject)
		projects.PATCH("/update", updateProject)
		projects.DELETE("/delete", deleteProject)
	}
	modules := r.Group("/" + rModule)
	{
		modules.GET("/list", listModules)
		modules.GET("/get", findModuleById)
		modules.GET("/get/:"+rModule, findModuleByName)
		modules.POST("/add", addModule)
		modules.PATCH("/update", updateModule)
		modules.DELETE("/delete", deleteModule)
	}
	//functions := r.Group("/" + rFunction)
	//{
	//	functions.GET("/list", listModules)
	//	functions.GET("/get", findModuleById)
	//	functions.GET("/get/:"+rModule, findModuleByName)
	//	functions.POST("/add", addModule)
	//	functions.PATCH("/update", updateModule)
	//	functions.DELETE("/delete", deleteModule)
	//}
	//	types := projects.Group("/" + rType)
	//	{
	//		types.GET("/", func(c *gin.Context) {
	//			c.String(http.StatusOK, "list of all types in project "+c.Param(rProject))
	//		})
	//		types.GET("/list/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "list of all types in module "+c.Param(rModule)+" from project"+c.Param(rProject))
	//		})
	//		//modules.POST("/add/:" + rModule, func(c *gin.Context) {
	//		types.GET("/add/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Adding type in module "+c.Param(rModule)+" from project "+c.Param(rProject))
	//		})
	//		//modules.PATCH("/update/:" + rName, func(c *gin.Context) {
	//		types.GET("/update/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Updating type in module "+c.Param(rModule)+" from project "+c.Param(rProject))
	//		})
	//		//modules.DELETE("/del/:" + rName, func(c *gin.Context) {
	//		types.GET("/del/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Deleting type in module "+c.Param(rModule)+" from project "+c.Param(rProject))
	//		})
	//	}
}
