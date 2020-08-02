package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // need for gorm
	"log"
	"os"
)

// DB is the database pointer
type Database struct {
	DB *gorm.DB
}

// Init : Initialise the db
func Init() (db *Database, err error) {
	db = &Database{}
	if os.Getenv("GIN_MODE") == "release" {
		log.Print("Database in Production mode")
		db.DB, err = gorm.Open("sqlite3", "./prod.db")
	} else {
		log.Print("Database in Debug mode")
		db.DB, err = gorm.Open("sqlite3", "./dev.db")
	}
	if err != nil {
		return nil, err
	}
	return db, err
}

// Add Database to gin context
func SetDatabase(db *Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

// Destroy db
func (db *Database) Destroy() {
	err := db.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
