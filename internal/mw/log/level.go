package log

import (
	"github.com/grimerssy/go-example/pkg/log"
	"google.golang.org/grpc/codes"
)

func getLogLevel(code codes.Code) log.Level {
	m := map[codes.Code]log.Level{
		codes.OK:                 log.Info,
		codes.Canceled:           log.Info,
		codes.Unknown:            log.Error,
		codes.InvalidArgument:    log.Info,
		codes.DeadlineExceeded:   log.Warn,
		codes.NotFound:           log.Info,
		codes.AlreadyExists:      log.Info,
		codes.PermissionDenied:   log.Warn,
		codes.ResourceExhausted:  log.Warn,
		codes.FailedPrecondition: log.Warn,
		codes.Aborted:            log.Warn,
		codes.OutOfRange:         log.Warn,
		codes.Unimplemented:      log.Error,
		codes.Internal:           log.Error,
		codes.Unavailable:        log.Warn,
		codes.DataLoss:           log.Error,
		codes.Unauthenticated:    log.Info,
	}
	return m[code]
}
