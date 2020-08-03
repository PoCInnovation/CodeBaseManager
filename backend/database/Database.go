package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // need for gorm
	"log"
	"os"
)

// Database: Typedef For holding Generic database, DB is the database pointer
type Database struct {
	DB *gorm.DB
}

// BackendDB: Package variable of type Database for easy access across backend
var BackendDB = &Database{}

// Init : Initialise the Database Struct
func Init() (db *Database, err error) {
	if os.Getenv("GIN_MODE") == "release" {
		log.Print("Database in Production mode")
		BackendDB.DB, err = gorm.Open("sqlite3", "./prod.db")
	} else {
		log.Print("Database in Debug mode")
		BackendDB.DB, err = gorm.Open("sqlite3", "./dev.db")
	}
	if err != nil {
		return nil, err
	}
	return BackendDB, err
}

// Destroy : Close the Database
func (db *Database) Destroy() {
	err := db.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
