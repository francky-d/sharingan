package models

import (
	"gorm.io/gorm"
)

type ApplicationGroup struct {
	gorm.Model
	Name         uint
	UserID       uint
	User         User
	Applications []Application
}

func (group ApplicationGroup) Table() string {
	return "application_groups"
}
