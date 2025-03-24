package main

import (
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/migrations"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/routes"
)

func init() {
	migrations.Migrate()
}

func main() {
	routes.Run()
}
