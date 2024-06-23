package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bitverb/x/errors"
)

// BaseResponse is the base response struct.
type BaseResponse struct {
	ErrNo  int         `json:"err_no"`         // business err no
	ErrMsg string      `json:"err_msg"`        // business err msg
	Data   interface{} `json:"data,omitempty"` // response data
}

func OkJson[T any](ctx *gin.Context, v T) {
	ctx.AbortWithStatusJSON(wrapBaseResponse(v))
}

// Err 提示错误信息,err 是预先定义的错误提示代码, msg 是运行时的错误
func Err[T any](ctx *gin.Context, err error, msg T) {
	ctx.AbortWithStatusJSON(wrapBaseErr(err, msg))
}

func wrapBaseErr(err error, msg any) (statusCode int, resp BaseResponse) {
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

func wrapBaseResponse(v interface{}) (statusCode int, resp BaseResponse) {

	resp.ErrNo = BusinessCodeOK
	resp.ErrMsg = BusinessMsgOk
	resp.Data = v
	statusCode = http.StatusOK

	return statusCode, resp
}
