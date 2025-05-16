package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiSuccessResponse struct {
	success bool
	ctx     *gin.Context
	data    interface{}
}

func NewApiSuccessResponse() *ApiSuccessResponse {
	return &ApiSuccessResponse{success: true}
}

func (resp *ApiSuccessResponse) SetContext(ctx *gin.Context) {
	resp.ctx = ctx
}

func (resp *ApiSuccessResponse) SetData(data any) {
	resp.data = data
}

func (resp *ApiSuccessResponse) sendResponse(status int, data any) {
	resp.ctx.JSON(status, gin.H{
		"data": data,
	})
}

func (resp *ApiSuccessResponse) SendCreated(data interface{}) {
	resp.ctx.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func (resp *ApiSuccessResponse) SendOk(data interface{}) {
	resp.ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (resp *ApiSuccessResponse) SendNoContent() {
	resp.ctx.JSON(http.StatusNoContent, nil)
}
