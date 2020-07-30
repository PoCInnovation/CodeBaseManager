package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

const (
    rModule = "module"
    rName = "name"
)


func NewRouter() *gin.Engine {
    r := gin.Default()

    modules := r.Group("/:" + rModule + "/:" + rName)
    {
        modules.GET("/", func(c *gin.Context) {
            c.String(http.StatusOK, "Welcome to " + c.Param(rModule))
        })
        modules.GET("/list", GH)
    }
    return r
}
