package models

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Path   string `gorm:"size:255;not null" json:"path"`
	Module Module `json:"module"`
}
