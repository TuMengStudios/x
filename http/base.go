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

// Err 提示错误信息,err 是预先定义的错误提示代码, msg 是运行时的错误
func Err[T any](ctx *gin.Context, err error, msg T) {
	ctx.AbortWithStatusJSON(wrapBaseErr[T](err, msg))
}

func wrapBaseErr[T any](err error, msg T) (statusCode int, resp BaseResponse[any]) {
	switch data := err.(type) {
	case *errors.CodeMsg:
		resp.ErrNo = data.ErrNo
		resp.ErrMsg = data.ErrMsg
		resp.Data = msg
		statusCode = data.StatusCode

	default:
		resp.ErrNo = BusinessCodeError
		resp.ErrMsg = data.Error()
		resp.Data = msg
		statusCode = http.StatusOK
	}

	return statusCode, resp
}

func wrapBaseResponse(v any) (statusCode int, resp BaseResponse[any]) {

	resp.ErrNo = BusinessCodeOK
	resp.ErrMsg = BusinessMsgOk
	resp.Data = v
	statusCode = http.StatusOK

	return statusCode, resp
}
