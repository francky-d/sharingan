package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/api/docs"
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

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middlewares.AuthenticationMiddleware())
	applicationGroupRoutes(apiV1)

}

func ConstructRouter(logger *zap.Logger) *gin.Engine {
	router := gin.New()
	// Apply CORS middleware first, before any other middleware
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig), middlewares.RequestsHistoryMiddleware(logger), middlewares.GinRecoverMiddleWare(logger))
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	logger.Debug("YOYO")
	docs.SwaggerInfo.BasePath = "/api/v1"

	constructRoutes(router)
	return router
}
