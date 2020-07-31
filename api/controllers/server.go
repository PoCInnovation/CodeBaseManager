package controllers

import (
	"cbm-api/database"
	"cbm-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Port   string
	Router *gin.Engine
	//Router *mux.Router
	//DB database.Database
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

	if err := database.CbmDb.Init(); err != nil {
		log.Fatalf("Database Initialisation Failed: %v", err)
	}
	s.Router = routes.NewRouter()
}

func (s *Server) Destroy() {
	database.CbmDb.Destroy()
}

func (s *Server) HandelerCores() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))
}
