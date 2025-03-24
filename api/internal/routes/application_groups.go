package routes

import "github.com/gin-gonic/gin"

func applicationGroupRoutes(apiV1 *gin.RouterGroup) {
	appGrps := apiV1.Group("/applications-groups")

	appGrps.GET("/", applictionGrpController.Index)
	appGrps.POST("/", applictionGrpController.Store)
	appGrps.GET("/:id", applictionGrpController.Show)
	appGrps.PUT(":id/update", applictionGrpController.Update)
	appGrps.DELETE(":id/delete", applictionGrpController.Delete)
}
