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

func NewServer() (*Server, func()) {
	server := &Server{}
	server.Init()
	return server, server.Destroy
}

func (s *Server) Init() {
	s.Port = os.Getenv("API_PORT")
	if s.Port == "" {
		s.Port = "5342"
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

	proj := s.Router.PathPrefix("/project/").Subrouter()
	proj.HandleFunc("/{project_name}/", middlewares.SetMiddlewareJSON(Hello)).Methods("GET")
	proj.HandleFunc("/list", middlewares.SetMiddlewareJSON(Hello)).Methods("GET")
	proj.HandleFunc("/add", middlewares.SetMiddlewareJSON(s.CreateProject)).Methods("POST")
	proj.HandleFunc("/{project_name}/", middlewares.SetMiddlewareJSON(Hello)).Methods("PUT")
	proj.HandleFunc("/{project_name}/", middlewares.SetMiddlewareJSON(Hello)).Methods("DELETE")

	mod := s.Router.PathPrefix("/module").Subrouter()
	mod.HandleFunc("/{project_name}/{module_id}/", middlewares.SetMiddlewareJSON(Hello)).Methods("GET")
	mod.HandleFunc("/{project_name}/add/", middlewares.SetMiddlewareJSON(Hello)).Methods("POST")
	mod.HandleFunc("/{project_name}/list/", middlewares.SetMiddlewareJSON(Hello)).Methods("GET")
	mod.HandleFunc("/{project_name}/{module_id}/", middlewares.SetMiddlewareJSON(Hello)).Methods("PUT")
	mod.HandleFunc("/{project_name}/{module_id}/", middlewares.SetMiddlewareJSON(Hello)).Methods("DEL")

	file := s.Router.PathPrefix("/file").Subrouter()
	file.HandleFunc("/{project_name}/{module_id}/", middlewares.SetMiddlewareJSON(Hello)).Methods("GET")
}
