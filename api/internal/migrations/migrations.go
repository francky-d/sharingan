package migrations

import (
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/contract"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/models"
)

type TableMappingToModel map[string]contract.ModelInterface

func Migrations() TableMappingToModel {

	return TableMappingToModel{
		"application_groups": &models.ApplicationGroup{},
		"applications":       &models.Application{},
		"incidents":          &models.Incident{},
	}
}
