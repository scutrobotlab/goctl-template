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

// CodeMsgResp is the base response struct.
type CodeMsgResp struct {
	// Code represents the business code, not the http status code.
	Code int `json:"code" xml:"code"`
	// Msg represents the business message, if Code = BusinessCodeOK,
	// and Msg is empty, then the Msg will be set to BusinessMsgOk.
	Msg string `json:"msg" xml:"msg"`
}

// JsonBaseResponseCtx writes v into w with http.StatusOK.
func JsonBaseResponseCtx(ctx context.Context, w http.ResponseWriter, v any) {
	if err, ok := v.(error); ok {
		code, resp := wrapError(err)
		httpx.WriteJsonCtx(ctx, w, code, resp)
	} else {
		httpx.WriteJsonCtx(ctx, w, http.StatusOK, v)
	}
}

func wrapError(err error) (int, CodeMsgResp) {
	var resp CodeMsgResp
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

	return runtime.HTTPStatusFromCode(codes.Code(resp.Code)), resp
}
