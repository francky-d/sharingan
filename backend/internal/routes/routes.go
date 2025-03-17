package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	constructRoutes(router)
	router.Run(":8000")
}
