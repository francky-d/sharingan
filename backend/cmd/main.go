package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/migrations"
)

func init() {
	migrations.Migrate()
}

func main() {

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.Run(":8000")
}
