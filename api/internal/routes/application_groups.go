package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func applicationGroupRoutes(apiV1 *gin.RouterGroup) {
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
