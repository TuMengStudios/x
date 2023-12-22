package http

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"

	"github.com/bitverb/x/errors"
)

// BaseResponse is the base response struct.
type BaseResponse[T any] struct {
	ErrNo  int    `json:"err_no"`         // business err no
	ErrMsg string `json:"err_msg"`        // business err msg
	Data   T      `json:"data,omitempty"` // response data
}

func OkJson(w http.ResponseWriter, v any) {
	httpx.OkJson(w, wrapBaseResponse(v))
}

func JsonErr(w http.ResponseWriter, err error) {
	httpx.OkJson(w, wrapBaseResponse(err))
}

func JsonErrCtx(ctx context.Context, w http.ResponseWriter, v error) {
	httpx.OkJsonCtx(ctx, w, wrapBaseResponse(v))
}

func OkJsonCtx(ctx context.Context, w http.ResponseWriter, v any) {
	httpx.OkJsonCtx(ctx, w, wrapBaseResponse(v))
}

func wrapBaseResponse(v any) BaseResponse[any] {
	var resp BaseResponse[any]
	switch data := v.(type) {
	case *errors.CodeMsg:
		resp.ErrNo = data.ErrNo
		resp.ErrMsg = data.ErrMsg
	case errors.CodeMsg:
		resp.ErrNo = data.ErrNo
		resp.ErrMsg = data.ErrMsg
	case *status.Status:
		resp.ErrNo = int(data.Code())
		resp.ErrMsg = data.Message()
	case error:
		resp.ErrNo = BusinessCodeError
		resp.ErrMsg = data.Error()
	default:
		resp.ErrNo = BusinessCodeOK
		resp.ErrMsg = BusinessMsgOk
		resp.Data = v
	}

	return resp
}
