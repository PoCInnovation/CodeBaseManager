package models

import (
	"cbm-api/database"
)

func MigrateModels() {
	database.BackendDB.DB.AutoMigrate(
		&Project{},
		&Module{},
		&Function{},
		&Type{},
		&Todo{},
	)
}
