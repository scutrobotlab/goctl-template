package http

import (
	"context"
	"errors"

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

	var statusCode int
	switch codes.Code(resp.Code) {
	case codes.OK: // 0
		statusCode = http.StatusOK // 200
	case codes.Unknown: // 2
		statusCode = http.StatusInternalServerError // 500
	case codes.InvalidArgument: // 3
		statusCode = http.StatusBadRequest // 400
	case codes.NotFound: // 5
		statusCode = http.StatusNotFound // 404
	case codes.PermissionDenied: // 7
		statusCode = http.StatusForbidden // 403
	case codes.FailedPrecondition: // 9
		statusCode = http.StatusPreconditionFailed // 412
	case codes.Unimplemented: // 12
		statusCode = http.StatusNotImplemented // 501
	case codes.Internal: // 13
		statusCode = http.StatusInternalServerError // 500
	case codes.Unavailable: // 14
		statusCode = http.StatusServiceUnavailable // 503
	case codes.DataLoss: // 15
		statusCode = http.StatusInternalServerError // 500
	case codes.Unauthenticated: // 16
		statusCode = http.StatusUnauthorized // 401
	case
		codes.Canceled,          // 1
		codes.DeadlineExceeded,  // 4
		codes.AlreadyExists,     // 6
		codes.ResourceExhausted, // 8
		codes.Aborted,           // 10
		codes.OutOfRange:        // 11
		statusCode = http.StatusInternalServerError // 500
	default:
		statusCode = http.StatusOK // 200
	}

	return statusCode, resp
}
