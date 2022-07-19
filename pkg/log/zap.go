package log

import (
	"time"

	"go.uber.org/zap"
)

type Zap struct {
	logger zap.Logger
}

func (z Zap) Log(lvl Level, msg string, fields ...Field) {
	z.getLogFunc(lvl)(msg, z.convertFields(fields)...)
}

func (Zap) WithString(key, val string) Field {
	return Field{
		zap: zap.String(key, val),
	}
}

func (Zap) WithStrings(key string, val []string) Field {
	return Field{
		zap: zap.Strings(key, val),
	}
}

func (Zap) WithDuration(key string, val time.Duration) Field {
	return Field{
		zap: zap.Duration(key, val),
	}
}

func (z Zap) getLogFunc(lvl Level) func(msg string, fields ...zap.Field) {
	switch lvl {
	case Debug:
		return z.logger.Debug
	case Info:
		return z.logger.Info
	case Warn:
		return z.logger.Warn
	case Error:
		return z.logger.Error
	case Fatal:
		return z.logger.Fatal
	case Panic:
		return z.logger.Panic
	default:
		panic("did not match the log level")
	}
}

func (Zap) convertFields(fs []Field) []zap.Field {
	zfs := make([]zap.Field, len(fs))
	for i, v := range fs {
		zfs[i] = v.zap
	}
	return zfs
}
