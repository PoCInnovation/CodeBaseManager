package codebase

import (
	"cbm-api/models"
	"github.com/jinzhu/gorm"
)

type Type struct {
	gorm.Model
	Path   string        `gorm:"size:255;not null" json:"path"`
	Module models.Module `json:"module"`
}
