package models

type Type struct {
	Path string `gorm:"size:255;not null" json:"path"`
}
