package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // need for gorm
)

// DB is the database pointer
type Database struct {
	DB *gorm.DB
}

// Init : Initialise the db
func (db *Database) Init() error {
	var err error
	db.DB, err = gorm.Open("sqlite3", "./dev.db")
	if err != nil {
		return err
	}
	//DB.CreateTable(&model.Users{})
	return err
}

// Destroy d
func (db *Database) Destroy() {
	db.DB.Close()
}
