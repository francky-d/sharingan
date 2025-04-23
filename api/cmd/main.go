package main

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/migrations"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/routes"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading environments variables %v", err)
	}

	migrations.Migrate()
}

func main() {
	routes.Run()
}
