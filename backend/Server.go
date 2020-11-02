package main

import (
	"github.com/PoCInnovation/CodeBaseManager/backend/database"
	"github.com/PoCInnovation/CodeBaseManager/backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Server struct {
	Port   string
	Router *gin.Engine
	Db     *database.Database
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

	var err error
	if s.Db, err = database.Init(); err != nil {
		log.Fatalf("Database Initialisation Failed: %v", err)
	}
	model.MigrateModels()

	s.Router = gin.Default()
}

func (s *Server) Destroy() {
	s.Db.Destroy()
}
