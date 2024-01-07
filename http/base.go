package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bitverb/x/errors"
)

// BaseResponse is the base response struct.
type BaseResponse[T any] struct {
	ErrNo  int    `json:"err_no"`         // business err no
	ErrMsg string `json:"err_msg"`        // business err msg
	Data   T      `json:"data,omitempty"` // response data
}

func OkJson(ctx *gin.Context, v any) {
	ctx.AbortWithStatusJSON(wrapBaseResponse(v))
}

func Err(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(wrapBaseResponse(err))
}

func wrapBaseResponse(v any) (statusCode int, resp BaseResponse[any]) {
	switch data := v.(type) {
	case *errors.CodeMsg:
		resp.ErrNo = data.ErrNo
		resp.ErrMsg = data.ErrMsg
		statusCode = data.StatusCode

	case errors.CodeMsg:
		resp.ErrNo = data.ErrNo
		resp.ErrMsg = data.ErrMsg
		statusCode = data.StatusCode

	case error:
		resp.ErrNo = BusinessCodeError
		resp.ErrMsg = data.Error()
		statusCode = http.StatusOK

	default:
		resp.ErrNo = BusinessCodeOK
		resp.ErrMsg = BusinessMsgOk
		resp.Data = v
		statusCode = http.StatusOK
	}

	return statusCode, resp
}
