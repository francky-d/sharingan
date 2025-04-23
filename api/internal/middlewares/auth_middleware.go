package middlewares

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	customErrors "gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/custom-errors"
	"net/http"
	"os"
	"strings"
	"time"
)

type KeycloakConfig struct {
	Host         string
	Realm        string
	ClientID     string
	ClientSecret string
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := getTokenFromRequest(c)
		keycloakConfig := getKeycloakConfig()
		keycloakClient := getKeycloakClient(keycloakConfig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if errors.Is(err, customErrors.TokenNotPresentErr) || errors.Is(err, customErrors.MustBeBearerToken) {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors":  err.Error(),
			})

			c.Abort()
			return
		}

		result, err := verifyTokenAgainstKeycloak(keycloakClient, keycloakConfig, accessToken, ctx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errors":  customErrors.InternalServerErr.Error(),
			})
			c.Abort()
			return
			//TODO:  logError
		}

		if !*result.Active {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"errors":  customErrors.UnauthorizedErr.Error(),
			})

			c.Abort()
			return
		}

		user, err := fetchAuthenticatedUserFromKeycloak(keycloakClient, ctx, accessToken, keycloakConfig.Realm)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errors":  "An error occurred while retrieving user informations",
			})
			//TODO:  logError
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()

		return
	}
}

func verifyTokenAgainstKeycloak(keycloakClient *gocloak.GoCloak, keycloakConfig KeycloakConfig, accessToken string, ctx context.Context) (*gocloak.IntroSpectTokenResult, error) {

	result, err := keycloakClient.RetrospectToken(ctx, accessToken, keycloakConfig.ClientID, keycloakConfig.ClientSecret, keycloakConfig.Realm)

	return result, err
}

func fetchAuthenticatedUserFromKeycloak(client *gocloak.GoCloak, ctx context.Context, accessToken string, keycloakRealm string) (*gocloak.UserInfo, error) {
	userInfo, err := client.GetUserInfo(ctx, accessToken, keycloakRealm)
	return userInfo, err
}

func getTokenFromRequest(c *gin.Context) (string, error) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return "", customErrors.TokenNotPresentErr
	}

	accessToken, isBearerToken := checkIfItsBearerToken(token)

	if !isBearerToken {
		return "", customErrors.MustBeBearerToken
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

func getKeycloakClient(keycloakConfig KeycloakConfig) *gocloak.GoCloak {
	return gocloak.NewClient(keycloakConfig.Host)
}

func getKeycloakConfig() KeycloakConfig {
	return KeycloakConfig{
		Host:         "http://keycloak:8080",
		Realm:        "sharingan",
		ClientID:     "sharingan-api",
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
	}
}
