package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Path string `gorm:"size:255;not null" json:"path"`
}
