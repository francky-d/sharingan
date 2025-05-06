package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/controllers"
	"go.uber.org/zap"
	"net/http"
)

func applicationGroupRoutes(apiV1 *gin.RouterGroup, logger *zap.Logger) {
	var applicationGrpController = controllers.NewApplicationGroupController(logger)

	appGrps := apiV1.Group("/applications-groups")
	appGrps.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/applications-groups/")
	})

	appGrps.GET("/", applicationGrpController.Index)
	appGrps.POST("/", applicationGrpController.Store)
	appGrps.GET("/:id", applicationGrpController.Show)
	appGrps.PUT(":id/update", applicationGrpController.Update)
	appGrps.DELETE(":id/delete", applicationGrpController.Delete)
}
