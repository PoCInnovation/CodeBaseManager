package model

import (
	"cbm-api/database"
)

// MigrateModels: Use gorm to make database.Database necessary migrations.
func MigrateModels() {
	database.BackendDB.DB.AutoMigrate(
		&Project{},
		&Module{},
		&Function{},
		&Type{},
		&Todo{},
	)
}
