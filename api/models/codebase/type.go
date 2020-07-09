package codebase

import "github.com/jinzhu/gorm"

type Type struct {
	gorm.Model
	Path string `gorm:"size:255;not null" json:"path"`
}
