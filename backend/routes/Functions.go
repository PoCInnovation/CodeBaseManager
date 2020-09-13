package routes

import (
	"github.com/PoCFrance/CodeBaseManager/backend/controllers"
	"github.com/PoCFrance/CodeBaseManager/backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// addFunction: Call controllers.AddFunction with moduleId, functionName and functionPath Query params.
// Abort when error occurs.
func addFunction(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryModule.ID = uint(moduleId)

	queryFunction := &model.Function{
		Name: c.Query("functionName"),
		Path: c.Query("functionPath"),
	}

	if function, err := controllers.AddFunction(queryModule, queryFunction); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusCreated, function)
	}
}

// listFunctions: Call controllers.ListFunctions with moduleId Query param.
// Abort when error occurs.
func listFunctions(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryModule.ID = uint(moduleId)

	if functions, err := controllers.ListFunctions(queryModule); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, functions)
	}
}

// findFunctionById: Call controllers.FindFunctionById with Query param.
// Abort when error occurs.
func findFunctionById(c *gin.Context) {
	queryFunction := &model.Function{}

	functionId, err := strconv.ParseInt(c.Query("functionId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryFunction.ID = uint(functionId)

	if function, err := controllers.FindFunctionById(queryFunction); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, function)
	}
}

// findFunctionByName: Call controllers.FindFunctionByName with module name in URL param and "projectId" in Query param.
// Abort when error occurs.
func findFunctionByName(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryModule.ID = uint(moduleId)

	queryFunction := &model.Function{
		Name: c.Param(rFunction),
	}

	if functions, err := controllers.FindFunctionByName(queryModule, queryFunction); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, functions)
	}
}

// findFunctionByPath: Call controllers.FindFunctionByName with module name in URL param and "projectId" in Query param.
// Abort when error occurs.
func findFunctionByPath(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("moduleId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryModule.ID = uint(moduleId)

	queryFunction := &model.Function{
		Name: c.Query("functionPath"),
	}

	if functions, err := controllers.FindFunctionByPath(queryModule, queryFunction); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, functions)
	}
}

// updateFunction: Call controllers.UpdateFunction with functionId Query param.
// Abort when error occurs.
func updateFunction(c *gin.Context) {
	queryFunction := &model.Function{}

	functionId, err := strconv.ParseInt(c.Query("functionId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryFunction.ID = uint(functionId)

	updatedFields := &model.Function{
		Name: c.Query("functionName"),
		Path: c.Query("functionPath"),
	}

	if function, err := controllers.UpdateFunction(queryFunction, updatedFields); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, function)
	}
}

// deleteFunction: Call controllers.DeleteFunction with functionId Query param.
// Abort when error occurs.
func deleteFunction(c *gin.Context) {
	queryFunction := &model.Function{}

	functionId, err := strconv.ParseInt(c.Query("functionId"), 10, 64)
	if err != nil {
		_ = c.AbortWithError(InternalError, err)
	}
	queryFunction.ID = uint(functionId)
	if function, err := controllers.DeleteFunction(queryFunction); err != nil {
		_ = c.AbortWithError(InternalError, err)
	} else {
		c.JSON(http.StatusOK, function)
	}
}
