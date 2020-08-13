package models

import (
	"github.com/jinzhu/gorm"
)

type Module struct {
	gorm.Model
	Name    string  `gorm:"size:255;not null" json:"name"`
	Path    string  `gorm:"size:255;not null" json:"path"`
	Project Project `json:"project"`
	//ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	//CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	//UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
