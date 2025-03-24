package models

import (
	"gorm.io/gorm"
)

type ApplicationGroup struct {
	gorm.Model
	Name         string        `json:"name" binding:"required,min=1,max=255"`
	UserID       uint          `json:"user_id"`
	User         User          `json:"user"`
	Applications []Application `json:"applications"`
}

func (group ApplicationGroup) Table() string {
	return "application_groups"
}
