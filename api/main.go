package main

import (
	"cbm-api/controllers"
	"cbm-api/routes"
	"log"
)

//func setDatabase(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Set("db", db)
//		c.Next()
//	}
//}
//server.Router.Use(database(db))

//Start Server -> Serve routes -> Defer server destroy
func main() {
	// TODO: add go routine for watcher.
	// TODO: Find how to request to api
	// TODO: change database management with above function call
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
