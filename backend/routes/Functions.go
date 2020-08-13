package routes

import (
	"cbm-api/controllers"
	"cbm-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func findFunctionByName(c *gin.Context) {
	queryModule := &model.Module{}

	moduleId, err := strconv.ParseInt(c.Query("functionId"), 10, 64)
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
