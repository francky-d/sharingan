package custom_errors

import "errors"

var TokenNotPresentErr = errors.New("token not present")
var MustBeBearerToken = errors.New("must be a bearer token")
var InternalServerErr = errors.New("something went wrong.Please,Try later")
var UnauthorizedErr = errors.New("unauthorized")
var OnGettingAuthenticatedUser = errors.New("error while getting user from context")
