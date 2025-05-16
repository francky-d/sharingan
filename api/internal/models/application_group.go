package models

import (
	"gorm.io/gorm"
)

type ApplicationGroup struct {
	gorm.Model
	UserID       string        `json:"user_id" gorm:"foreignKey:-"`
	Name         string        `json:"name" binding:"required,min=1,max=255"`
	Applications []Application `json:"applications"`
}

func (group ApplicationGroup) Table() string {
	return "application_groups"
}
