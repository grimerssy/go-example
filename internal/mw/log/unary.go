package log

import (
	"context"

	"github.com/grimerssy/go-example/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func UnaryServerInterceptor(logger logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		stopTimer := startTimer()
		res, err := handler(ctx, req)
		code := getGrpcCode(err)
		logger.Log(
			getLogLevel(code),
			err.Error(),
			logger.WithString("gRPC-code", code.String()),
			logger.WithStrings("callers", getCallers(err)),
			logger.WithDuration("time-taken", stopTimer()))
		return res, err
	}
}

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

func getGrpcCode(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	c, ok := err.(interface{ Code() codes.Code })
	if !ok {
		return codes.Unknown
	}
	return c.Code()
}

func getCallers(err error) []string {
	c, ok := err.(interface{ Callers() []string })
	if !ok {
		return []string{}
	}
	return c.Callers()
}
