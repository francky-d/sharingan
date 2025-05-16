package response

import "github.com/gin-gonic/gin"

type Response struct {
	error   *ApiErrorResponse
	success *ApiSuccessResponse
	ctx     *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		error:   NewApiErrorResponse(),
		success: NewApiSuccessResponse(),
		ctx:     ctx,
	}
}

func (response *Response) Error() *ApiErrorResponse {
	response.error.SetContext(response.ctx)
	return response.error
}

func (response *Response) Success() *ApiSuccessResponse {
	response.success.SetContext(response.ctx)
	return response.success
}
