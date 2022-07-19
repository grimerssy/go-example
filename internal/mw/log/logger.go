package log

import (
	"time"

	"github.com/grimerssy/go-example/pkg/log"
)

type logger interface {
	Log(lvl log.Level, msg string, fields ...log.Field)
	WithString(key, val string) log.Field
	WithStrings(key string, ss []string) log.Field
	WithDuration(key string, val time.Duration) log.Field
}
