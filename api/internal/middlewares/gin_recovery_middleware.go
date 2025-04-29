package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinRecoverMiddleWare(logger *zap.Logger) gin.HandlerFunc {
	return gin.RecoveryWithWriter(zap.NewStdLog(logger).Writer())
}
