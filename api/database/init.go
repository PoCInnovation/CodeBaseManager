package database

import (
	"cbm-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // need for gorm
)

// DB is the database pointer
type Database struct {
	DB *gorm.DB
}

func (db *Database) initTables() {
	db.DB.CreateTable(&models.Project{})
}

// Init : Initialise the db
func (db *Database) Init() (err error) {
	db.DB, err = gorm.Open("sqlite3", "./dev.db")
	if err != nil {
		return err
	}
	//DB.CreateTable(&model.Users{})
	db.DB.LogMode(true)
	db.initTables()
	db.DB.AutoMigrate()
	return err
}

// Destroy d
func (db *Database) Destroy() {
	db.DB.Close()
}
