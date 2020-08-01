package main

import (
	"cbm-api/controllers"
	"cbm-api/routes"
	"log"
)

//Start Server -> Serve routes -> Defer server destroy
func main() {
	// TODO: add go routine for watcher.
	// TODO: Find how to request to api
	server, closer := controllers.NewServer()

	routes.ApplyRoutes(server.Router)
	defer closer()
	log.Println("Server runs on http://localhost:" + server.Port)
	//if err := http.ListenAndServe(server.Port, server.Router); err != nil {
	//	log.Fatal(err)
	//}
	if err := server.Router.Run(); err != nil {
		log.Fatal(err)
	}
}
