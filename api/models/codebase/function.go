package codebase

import (
	"cbm-api/models"
	"github.com/jinzhu/gorm"
)

type Function struct {
	gorm.Model
	Path   string        `gorm:"size:255;not null" json:"path"`
	Type   Type          `json:"type"`
	Module models.Module `json:"module"`
}
