package controllers

import (
	"cbm-api/database"
	"cbm-api/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Port   string
	Router *mux.Router
	DB     database.Database
}

func (s *Server) Init() {
	s.Port = os.Getenv("API_PORT")
	if s.Port == "" {
		s.Port = "8080"
		log.Printf("Defaulting to port %s", s.Port)
	}

	if err := s.DB.Init(); err != nil {
		log.Fatalf("Database Initialisation Failed: %v", err)
	}

	s.Router = mux.NewRouter().StrictSlash(false)
	s.initialiseRoutes()
}

func (s *Server) Destroy() {
	s.DB.Destroy()
}

func (s *Server) HandelerCores() func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))
}

func (s *Server) initialiseRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(Home)).Methods("GET")
	s.Router.HandleFunc("/hello", middlewares.SetMiddlewareJSON(Hello)).Methods("GET")

}
