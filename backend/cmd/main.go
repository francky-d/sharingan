package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/database"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/migrations"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/models"
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
	appGrps := apiV1.Group("/applications-groups")

	appGrps.GET("/", func(c *gin.Context) {
		var applicationsGroups []models.ApplicationGroup

		database.DbConnection().Db().Find(&applicationsGroups)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"application_groups": applicationsGroups,
			},
		})
	})

	appGrps.POST("/", func(c *gin.Context) {
		var group models.ApplicationGroup

		if err := c.ShouldBindJSON(&group); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		group.UserID = uint(1)

		result := database.DbConnection().Db().Create(&group)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, result.Error.Error())
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data": gin.H{
				"id": group.ID,
			}},
		)

	})

	router.Run(":8000")
}
