package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/controllers"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/migrations"
)

func init() {
	migrations.Migrate()
}

func main() {

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": true,
		})
	})

	apiV1 := router.Group("/api/v1")
	applictionGrpController := controllers.NewApplicationGroupController()
	appGrps := apiV1.Group("/applications-groups")

	appGrps.GET("/", applictionGrpController.Index)

	appGrps.POST("/", applictionGrpController.Store)

	appGrps.GET("/:id", applictionGrpController.Show)

	appGrps.PUT(":id/update", applictionGrpController.Update)

	appGrps.DELETE(":id/delete", applictionGrpController.Delete)

	router.Run(":8000")
}
