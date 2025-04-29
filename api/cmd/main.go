package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/migrations"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/routes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading environments variables %v", err)
	}
	migrations.Migrate()
}

func NewLoggerWithFile() (*zap.Logger, error) {
	today := time.Now().Format("2006-01-02")

	err := os.MkdirAll("logs", 777)

	if err != nil {
		return nil, fmt.Errorf("error while creating/opening 'logs' directory %w", err)
	}

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("logs/gin_%s.log", today),
		MaxSize:    10, //megabytes
		MaxBackups: 3,
		MaxAge:     7, // days,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	logLevel := zap.InfoLevel
	if os.Getenv("APP_ENV") == "development" {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		logLevel = zap.DebugLevel
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		logLevel,
	)

	return zap.New(core), err
}

func main() {
	logger, err := NewLoggerWithFile()
	if err != nil {
		log.Fatalf("error while creating error : %v", err)
	}
	logger.Info("HELLOOOOOO")

	defer logger.Sync()
	router := routes.ConstructRouter(logger)

	logger.Fatal(router.Run(":" + os.Getenv("APP_PORT")).Error())
}
