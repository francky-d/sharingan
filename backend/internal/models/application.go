package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	UserID             uint
	ApplicationGroupID uint
	Url                string
	UrlToWatch         string
	HttpSuccessCode    int
	User               User
	ApplicationGroup   ApplicationGroup
	Incidents          []Incident
}

func (app Application) Table() string {
	return "applications"
}
