package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/api/docs"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/controllers"
)

var applictionGrpController = controllers.NewApplicationGroupController()

func constructRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": true,
		})
	})

	apiV1 := router.Group("/api/v1")

	applicationGroupRoutes(apiV1)

}

func Run() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	docs.SwaggerInfo.BasePath = "/api/v1"

	constructRoutes(router)
	router.Run(":8000")
}
