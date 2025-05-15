package custom_logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func NewLoggerWithFile() (*zap.Logger, error) {
	today := time.Now().Format("2006-01-02")

	err := os.MkdirAll("logs", 0755)

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
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		logLevel,
	)

	return zap.New(core), err
}
