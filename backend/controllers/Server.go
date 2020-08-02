package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Server struct {
	Port   string
	Router *gin.Engine
}

func NewServer() (*Server, func()) {
	server := &Server{}
	server.Init()
	return server, server.Destroy
}

func (s *Server) Init() {
	s.Port = os.Getenv("PORT")
	if s.Port == "" {
		s.Port = "5342"
		log.Printf("Defaulting to port %s", s.Port)
	}
	db, err := database.Init()
	if err != nil {
		log.Fatalf("Database Initialisation Failed: %v", err)
	}
	models.MigrateModels(db)
	s.Router = gin.Default()
	s.Router.Use(database.SetDatabase(db))
}

func (s *Server) Destroy() {
	database.CbmDb.Destroy()
}
