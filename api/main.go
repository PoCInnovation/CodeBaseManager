package main

import (
	"cbm-api/controllers"
	"log"
	"net/http"
)

//Start Server -> Serve routes -> Defer server destroy
func main() {
	server := new(controllers.Server)
	server.Init()

	defer server.Destroy()
	log.Println("Server runs on http://localhost:" + server.Port)
	log.Fatal(
		http.ListenAndServe(":"+server.Port, server.HandelerCores()(server.Router)))
}
