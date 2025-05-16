package models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	UserID             string `json:"user_id" gorm:"foreignKey:-"`
	ApplicationGroupID uint
	Url                string
	UrlToWatch         string
	HttpSuccessCode    int
	ApplicationGroup   ApplicationGroup
	Incidents          []Incident
}

func (app Application) Table() string {
	return "applications"
}
