package database

import (
	"cbm-api/models_v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // need for gorm
	"log"
	"os"
)

// DB is the database pointer
type Database struct {
	DB *gorm.DB
}

var CbmDb = Database{}

func (db *Database) initTables() {
	db.DB.CreateTable(&models_v2.Project{})
	db.DB.CreateTable(&models_v2.Module{})
	db.DB.CreateTable(&models_v2.Function{})
	db.DB.CreateTable(&models_v2.Type{})
}

// Init : Initialise the db
func (db *Database) Init() (err error) {
	if os.Getenv("PORT") == "release" {
		log.Print("Database in Production mode")
		db.DB, err = gorm.Open("sqlite3", "./prod.db")
	} else {
		log.Print("Database in Debug mode")
		db.DB, err = gorm.Open("sqlite3", "./dev.db")
	}
	if err != nil {
		return err
	}
	db.DB.LogMode(true)
	db.initTables()
	db.DB.AutoMigrate()
	return err
}

// Destroy d
func (db *Database) Destroy() {
	db.DB.Close()
}
