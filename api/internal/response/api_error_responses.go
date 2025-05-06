package response

import (
	"github.com/gin-gonic/gin"
	custom_errors "gitlab.jems-group.com/fdjacoto/sharingan/backend/internal/custom-errors"
	"net/http"
)

type ApiErrorResponse struct {
	success bool
	ctx     *gin.Context
	error   error
}

func NewApiErrorResponse() *ApiErrorResponse {
	return &ApiErrorResponse{
		success: false,
	}
}

func (resp *ApiErrorResponse) SetContext(ctx *gin.Context) {
	resp.ctx = ctx
}
func (resp *ApiErrorResponse) sendError(status int, err error) {
	resp.error = err
	resp.ctx.JSON(status, gin.H{
		"success": "false",
		"error":   resp.error.Error(),
	})
}

func (resp *ApiErrorResponse) SendBadRequestWithErr(err error) {
	resp.sendError(http.StatusBadRequest, err)
}

func (resp *ApiErrorResponse) SendUnauthorizedWithErr(err error) {
	resp.sendError(http.StatusUnauthorized, err)
}

func (resp *ApiErrorResponse) SendForbiddenWithErr(err error) {
	resp.sendError(http.StatusForbidden, err)
}

func (resp *ApiErrorResponse) SendInternalServerWithErr() {
	resp.sendError(http.StatusInternalServerError, custom_errors.InternalServerErr)
}

func (resp *ApiErrorResponse) SendNotFoundWithErr(err error) {
	resp.sendError(http.StatusNotFound, err)
}

func (resp *ApiErrorResponse) Abort() {
	resp.ctx.Abort()
}
