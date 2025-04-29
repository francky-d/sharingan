package core

import (
	"github.com/joho/godotenv"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/custom_logger"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/migrations"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/routes"
	"log"
	"os"
)

type App struct {
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading environments variables %v", err)
	}

	migrations.Migrate()
}

func (app App) Start() {
	logger, err := custom_logger.NewLoggerWithFile()
	if err != nil {
		log.Fatalf("error while creating logger : %v", err)
	}
	defer logger.Sync()

	router := routes.ConstructRouter(logger)

	logger.Fatal(router.Run(":" + os.Getenv("APP_PORT")).Error())
}
