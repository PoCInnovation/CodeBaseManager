package models

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Name    string `gorm:"size:255;not null" json:"name"`
	Project Project
	// TODO: add additional infos ??
	//ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	//CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	//UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
