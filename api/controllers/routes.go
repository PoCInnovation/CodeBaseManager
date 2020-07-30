package controllers

import (
	"cbm-api/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Home : Root endpoint
func Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "World")
}

func GHello(c *gin.Context) {
	c.String(http.StatusOK, "World")
}

