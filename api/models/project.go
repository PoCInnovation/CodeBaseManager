package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
	//ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	//CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	//UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Project) SaveProject(db *gorm.DB) (*Project, error) {
	var err error
	err = db.Create(&p).Error
	if err != nil {
		return &Project{}, err
	}
	return p, nil
}

func (p *Project) FindProjectByName(db *gorm.DB, name string) (*Project, error) {
	var err error
	err = db.Model(Project{}).Where("name = ?", name).Take(&p).Error
	if err != nil {
		return &Project{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Project{}, errors.New("user not found")
	}
	return p, err
}
