package models

import "gorm.io/gorm"

type Incident struct {
	gorm.Model
	UserID        uint
	ApplicationID uint
	Title         string
	Description   string
	User          User
	Application   Application
}

func (incident Incident) Table() string {
	return "incidents"
}
