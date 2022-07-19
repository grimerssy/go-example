package log

import (
	"github.com/grimerssy/go-example/pkg/log"
	"google.golang.org/grpc/codes"
)

func getLogLevel(code codes.Code) log.Level {
	switch code {
	case codes.OK:
		return log.Info
	case codes.Canceled:
		return log.Info
	case codes.Unknown:
		return log.Error
	case codes.InvalidArgument:
		return log.Info
	case codes.DeadlineExceeded:
		return log.Warn
	case codes.NotFound:
		return log.Info
	case codes.AlreadyExists:
		return log.Info
	case codes.PermissionDenied:
		return log.Warn
	case codes.ResourceExhausted:
		return log.Warn
	case codes.FailedPrecondition:
		return log.Warn
	case codes.Aborted:
		return log.Warn
	case codes.OutOfRange:
		return log.Warn
	case codes.Unimplemented:
		return log.Error
	case codes.Internal:
		return log.Error
	case codes.Unavailable:
		return log.Warn
	case codes.DataLoss:
		return log.Error
	case codes.Unauthenticated:
		return log.Info
	default:
		panic("did not match gRpc-code")
	}
}
