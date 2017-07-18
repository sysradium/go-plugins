package grpc

import (
	"google.golang.org/grpc/codes"
)

var hToGrpcCodes = map[int32]codes.Code{
	400: codes.InvalidArgument,
	401: codes.Unauthenticated,
	403: codes.Unauthenticated,
	404: codes.NotFound,
}

func convertHTTPCode(c int32) codes.Code {
	code, ok := hToGrpcCodes[c]
	if !ok {
		return codes.Internal
	}

	return code
}
