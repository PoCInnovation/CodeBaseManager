package routes

import (
	"github.com/gin-gonic/gin"
)

func addType(c *gin.Context) {
	//name := c.Query("name")
	//path := c.Query("path")
	//newType := model.Type{
	//	Name: name,
	//	Path: path,
	//}
	//var err error
	//if _, err = newType.SaveType(database.CbmDb.DB); err != nil {
	//	_ = c.AbortWithError(http.StatusForbidden, err)
	//} else {
	//	c.JSON(http.StatusCreated, newType)
	//}
}

func listType(c *gin.Context) {
	//result := database.CbmDb.DB.Find(&model.Type{})
	//if result.Error != nil {
	//	c.Value(http.StatusNotFound)
	//}
	//c.JSON(http.StatusOK, result)
	////c.String(http.StatusOK, "List of all project")
}

func findType(c *gin.Context) {
	//name := c.Param(rType)
	//toFindType := model.Type{
	//	Name: name,
	//}
	//
	//result := database.CbmDb.DB.First(&toFindType)
	//if result.Error != nil {
	//	c.Value(http.StatusNotFound)
	//}
	//c.JSON(http.StatusOK, result)
	//c.String(http.StatusOK, "List of all project")
}

func deleteType(c *gin.Context) {
	//name := c.Query("name")
	//path := c.Query("path")
	//
	//toFindType := model.Type{
	//	Name: name,
	//	Path: path,
	//}
	////var err error
	//
	//result := database.CbmDb.DB.First(&toFindType)
	//if result.Error != nil {
	//	c.Value(http.StatusNotFound)
	//	return
	//}
	//if _, err := toFindType.DeleteType(database.CbmDb.DB); err != nil {
	//	_ = c.AbortWithError(http.StatusNotFound, err)
	//} else {
	//	c.JSON(http.StatusOK, result)
	//}
}
