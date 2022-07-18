package log

import (
	"time"

	"github.com/grimerssy/go-example/pkg/log"
	"google.golang.org/grpc/codes"
)

type logger interface {
	Log(msg string, fields ...log.Field)
	WithGrpcCode(code codes.Code) log.Field
	WithStrings(key string, ss []string) log.Field
	WithDuration(key string, val time.Duration) log.Field
}
