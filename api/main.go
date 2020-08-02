package main

import (
	"cbm-api/controllers"
	"cbm-api/routes"
	"cbm-api/watcher"
	"log"
)

//func setDatabase(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Set("db", db)
//		c.Next()
//	}
//}
//server.Router.Use(database(db))

func main() {
	//Setup the watcher to keep track of projects' files & gets ready to properly close it
	stopWatcher := make(chan struct{})
	go watcher.Run(stopWatcher)
	defer close(stopWatcher)

	// Setup the server for CLI's needs & gets ready to properly close it
	server, stopServer := controllers.NewServer()
	// TODO: change database management with above function call
	routes.ApplyRoutes(server.Router)
	defer stopServer()

	// Starts the server
	log.Println("Server runs on http://localhost:" + server.Port)
	if err := server.Router.Run(); err != nil {
		log.Fatal(err)
	}
}
