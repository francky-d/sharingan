package models

import (
	"gorm.io/gorm"
)

type Incident struct {
	gorm.Model
	UserID        string `json:"user_id" gorm:"foreignKey:-"`
	ApplicationID uint
	Title         string
	Description   string
	Application   Application
}

func (incident Incident) Table() string {
	return "incidents"
}
