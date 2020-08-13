package main

import (
	"cbm-api/routes"
	"cbm-api/watcher"
	"log"
)

func main() {
	//Setup the watcher to keep track of projects' files & gets ready to properly close it
	stopWatcher := make(chan struct{})
	go watcher.Run(stopWatcher)
	defer close(stopWatcher)

	// Setup the server for CLI's needs & gets ready to properly close it
	server, stopServer := NewServer()
	routes.ApplyRoutes(server.Router)
	defer stopServer()

	// Starts the server
	log.Println("Server runs on http://localhost:" + server.Port)
	if err := server.Router.Run(); err != nil {
		log.Fatal(err)
	}
}
