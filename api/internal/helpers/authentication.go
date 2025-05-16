package helpers

import (
	"errors"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(ctx *gin.Context) (*gocloak.UserInfo, error) {

	value, userIsSet := ctx.Get("user")
	if !userIsSet {
		return nil, errors.New("authenticated user not set in context")
	}

	user, castIsSuccessful := value.(*gocloak.UserInfo)

	if !castIsSuccessful {
		return nil, errors.New("fail to cast user to *gocloak.UserInfo")
	}

	return user, nil
}
