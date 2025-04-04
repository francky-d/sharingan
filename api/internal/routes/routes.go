package routes

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	custom_errors "gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/custom-errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/api/docs"
	"gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/controllers"
)

func authenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := GetTokenFromRequest(c)
		if errors.Is(err, custom_errors.TokenNotPresentErr) || errors.Is(err, custom_errors.MustBeBearerToken) {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors":  err.Error(),
			})

			c.Abort()
			return
		}

		// TODO: get the config value from .env variables
		keycloackConfig := KeycloackConfig{
			Host:         "http://keycloak:8080",
			Realm:        "sharingan",
			ClientID:     "sharingan-api",
			ClientSecret: "k2ceR5OUCgjKEmbp4tHiUNxvRPF7EI7q",
		}
		client := gocloak.NewClient(keycloackConfig.Host)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		result, err := client.RetrospectToken(ctx, accessToken, keycloackConfig.ClientID, keycloackConfig.ClientSecret, keycloackConfig.Realm)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errors":  custom_errors.InternalServerErr.Error(),
			})
			c.Abort()
			return
			//TODO:  logError
		}

		if !*result.Active {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"errors":  custom_errors.UnauthorizedErr.Error(),
			})

			c.Abort()
			return
		}

		userInfo, err := client.GetUserInfo(ctx, accessToken, keycloackConfig.Realm)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errors":  "An error occurred while retrieving user informations",
			})
			//TODO:  logError
			c.Abort()
			return
		}

		c.Set("user", userInfo)

		c.JSON(200, gin.H{
			"success": "true",
			"data":    userInfo,
		})

		return
	}
}

func GetTokenFromRequest(c *gin.Context) (string, error) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return "", custom_errors.TokenNotPresentErr
	}

	accessToken, isBearerToken := checkIfItsBearerToken(token)

	if !isBearerToken {
		return "", custom_errors.MustBeBearerToken
	}
	return accessToken, nil

}

func checkIfItsBearerToken(token string) (string, bool) {
	parts := strings.Split(token, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1], true
	}
	return "", false
}

var applicationGrpController = controllers.NewApplicationGroupController()

func constructRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": true,
		})
	})

	apiV1 := router.Group("/api/v1")
	apiV1.Use(authenticationMiddleware())
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
