package log

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func UnaryServerInterceptor(logger logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		stopTimer := startTimer()
		res, err := handler(ctx, req)
		logger.Log(err.Error(),
			logger.WithGrpcCode(getGrpcCode(err)),
			logger.WithStrings("callers", getCallers(err)),
			logger.WithDuration("time-taken", stopTimer()))
		return res, err
	}
}

func getGrpcCode(err error) codes.Code {
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
