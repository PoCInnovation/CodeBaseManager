package models

import (
	"cbm-api/database"
)

func MigrateModels(db *database.Database) {
	db.DB.AutoMigrate(
		&Project{},
		&Module{},
		&Function{},
		&Type{},
		&Todo{},
	)
}
