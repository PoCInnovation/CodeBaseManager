package main

import (
	"cbm-api/routes"
	"log"
)

func main() {
	// Setup the server for CLI's needs & gets ready to properly close it
	server, stopServer := NewServer()
	routes.ApplyRoutes(server.Router)
	defer stopServer()

	// Starts the server
	log.Println("Server runs on http://localhost:" + server.Port)
	if err := server.Router.Run(); err != nil {
		log.Fatalln(err)
	}
}
