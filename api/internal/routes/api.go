package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/controllers"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/middlewares"
	"go.uber.org/zap"
	"net/http"
)

var applicationGrpController = controllers.NewApplicationGroupController()

func constructRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": true,
		})
	})

	swaggerRoutes(router)

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middlewares.AuthenticationMiddleware())

	applicationGroupRoutes(apiV1)

}

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
}

func applyGlobalMiddlewares(router *gin.Engine, logger *zap.Logger) {
	router.Use(cors.New(CorsConfig()), middlewares.RequestsHistoryMiddleware(logger), middlewares.GinRecoverMiddleWare(logger))

}

func Router(logger *zap.Logger) *gin.Engine {
	router := gin.New()
	applyGlobalMiddlewares(router, logger)
	constructRoutes(router)
	return router
}
