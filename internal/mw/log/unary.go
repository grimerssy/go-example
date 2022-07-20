package log

import (
	"context"

	"github.com/grimerssy/go-example/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor(logger Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {
		stopTimer := startTimer()
		res, err := handler(ctx, req)
		duration := stopTimer()
		code := status.Code(err)
		lvl := getLogLevel(code)
		msg := err.Error()
		fields := []log.Field{
			logger.WithString("gRPC-code", code.String()),
			logger.WithDuration("time-taken", duration),
		}
		callers := getCallers(err)
		if len(callers) != 0 {
			fields = append(fields, logger.WithStrings("callers", callers))
		}
		logger.Log(lvl, msg, fields...)
		return res, err
	}
}
