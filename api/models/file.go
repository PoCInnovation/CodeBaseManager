package models

type File struct {
	Path   string `gorm:"size:255;not null" json:"path"`
	Module Module `json:"module"`
}
