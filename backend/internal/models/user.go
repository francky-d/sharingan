package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Incidents    []Incident
	Groups       []ApplicationGroup
	Applications []Application
}

func (user User) Table() string {
	return "users"
}
