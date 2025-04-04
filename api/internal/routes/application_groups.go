package routes

import "github.com/gin-gonic/gin"

func applicationGroupRoutes(apiV1 *gin.RouterGroup) {
	appGrps := apiV1.Group("/applications-groups")

	appGrps.GET("/", applicationGrpController.Index)
	appGrps.POST("/", applicationGrpController.Store)
	appGrps.GET("/:id", applicationGrpController.Show)
	appGrps.PUT(":id/update", applicationGrpController.Update)
	appGrps.DELETE(":id/delete", applicationGrpController.Delete)
}
