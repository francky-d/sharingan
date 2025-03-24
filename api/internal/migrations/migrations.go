package migrations

import (
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/contract"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/models"
)

type tableMappingToModel map[string]contract.ModelInterface

func Migrations() tableMappingToModel {

	return tableMappingToModel{
		"users":              &models.User{},
		"application_groups": &models.ApplicationGroup{},
		"applications":       &models.Application{},
		"incidents":          &models.Incident{},
	}
}
