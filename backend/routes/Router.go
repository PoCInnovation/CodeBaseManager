package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	rModule   = "module"
	rProject  = "project"
	rFunction = "function"
	rType     = "type"
	rName     = "name"
)

func ApplyRoutes(r *gin.Engine) {

	projects := r.Group("/" + rProject)
	{
		projects.GET("/list", listProject)
		projects.GET("/get", findProject)
		projects.POST("/add", addProject)
		projects.DELETE("/delete", deleteProject)
	}
	modules := projects.Group("/" + rModule)
	{
		//modules.POST("/add", func(c *gin.Context) {
		modules.GET("/add", func(c *gin.Context) {
			c.String(http.StatusOK, "Adding module "+c.Param(rModule)+" in project "+c.Param(rProject))
		})
		//modules.PATCH("/update", func(c *gin.Context) {
		modules.GET("/update", func(c *gin.Context) {
			c.String(http.StatusOK, "Updating module "+c.Param(rModule)+" in project "+c.Param(rProject))
		})
		//modules.DELETE("/delete", func(c *gin.Context) {
		modules.GET("/delete", func(c *gin.Context) {
			c.String(http.StatusOK, "Deleting module "+c.Param(rModule)+" in project "+c.Param(rProject))
		})
	}
	//projects := r.Group("/project/:" + rProject)
	//{
	//	projects.GET("/", findProject)
	//	modules := projects.Group("/" + rModule)
	//	{
	//		//modules.POST("/add/:" + rName, func(c *gin.Context) {
	//		modules.GET("/add/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Adding module "+c.Param(rModule)+" in project "+c.Param(rProject))
	//		})
	//		//modules.PATCH("/update/:" + rName, func(c *gin.Context) {
	//		modules.GET("/update/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Updating module "+c.Param(rModule)+" in project "+c.Param(rProject))
	//		})
	//		//modules.DELETE("/del/:" + rName, func(c *gin.Context) {
	//		modules.GET("/del/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Deleting module "+c.Param(rModule)+" in project "+c.Param(rProject))
	//		})
	//	}
	//	function := projects.Group("/" + rFunction)
	//	{
	//		function.GET("/", func(c *gin.Context) {
	//			c.String(http.StatusOK, "list of all functions in project "+c.Param(rProject))
	//		})
	//		function.GET("/list/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "list of all functions in module "+c.Param(rModule)+" from project"+c.Param(rProject))
	//		})
	//		//modules.POST("/add/:" + rModule, func(c *gin.Context) {
	//		function.GET("/add/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Adding function in module "+c.Param(rModule)+" from project "+c.Param(rProject))
	//		})
	//		//modules.PATCH("/update/:" + rName, func(c *gin.Context) {
	//		function.GET("/update/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Updating function in module "+c.Param(rModule)+" from project "+c.Param(rProject))
	//		})
	//		//modules.DELETE("/del/:" + rName, func(c *gin.Context) {
	//		function.GET("/del/:"+rModule, func(c *gin.Context) {
	//			c.String(http.StatusOK, "Deleting function in module "+c.Param(rModule)+" from project "+c.Param(rProject))
	//		})
	//	}
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
	//}

	//modules := r.Group("/:" + rModule + "/:" + rFunction)
	//{
	//	modules.GET("/", func(c *gin.Context) {
	//		c.String(http.StatusOK, "Welcome to "+c.Param(rModule))
	//	})
	//	//modules.GET("/list", GH)
	//}
}
