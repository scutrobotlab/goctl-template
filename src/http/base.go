package http

import (
	"context"
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zeromicro/go-zero/rest/httpx"
	errors2 "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// BaseResponse is the base response struct.
type BaseResponse[T any] struct {
	// Code represents the business code, not the http status code.
	Code int `json:"code" xml:"code"`
	// Msg represents the business message, if Code = BusinessCodeOK,
	// and Msg is empty, then the Msg will be set to BusinessMsgOk.
	Msg string `json:"msg" xml:"msg"`
	// Data represents the business data.
	Data T `json:"data,omitempty" xml:"data,omitempty"`
}

// JsonBaseResponseCtx writes v into w with http.StatusOK.
func JsonBaseResponseCtx(ctx context.Context, w http.ResponseWriter, v any) {
	code, v := wrapBaseResponse(v)
	httpx.WriteJsonCtx(ctx, w, code, v)
}

func wrapBaseResponse(v any) (int, BaseResponse[any]) {
	var resp BaseResponse[any]
	if err, ok := v.(error); ok {
		var codeMsg *errors2.CodeMsg
		if errors.As(err, &codeMsg) {
			resp.Code = codeMsg.Code
			resp.Msg = codeMsg.Msg
		} else if grpcStatus, ok := status.FromError(err); ok {
			resp.Code = int(grpcStatus.Code())
			resp.Msg = grpcStatus.Message()
		} else {
			resp.Code = BusinessCodeError
			resp.Msg = err.Error()
		}
	} else {
		resp.Code = BusinessCodeOK
		resp.Msg = BusinessMsgOk
		resp.Data = v
	}

	return runtime.HTTPStatusFromCode(codes.Code(resp.Code)), resp
}
