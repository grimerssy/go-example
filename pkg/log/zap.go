package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ConfigZap struct {
	IsDevelopment bool
	Level         string
	DisableCaller bool
	Encoding      string
	OutputPaths   []string
}

type Zap struct {
	l *zap.Logger
}

func NewZap(cfg ConfigZap) *Zap {
	l := zap.NewProductionConfig()
	if cfg.IsDevelopment {
		l = zap.NewDevelopmentConfig()
	}
	l.Level = zap.NewAtomicLevelAt(toZapLevel(getLevel(cfg.Level)))
	l.DisableCaller = cfg.DisableCaller
	l.Encoding = cfg.Encoding
	l.OutputPaths = cfg.OutputPaths
	l.ErrorOutputPaths = cfg.OutputPaths
	logger, err := l.Build()
	if err != nil {
		panic(err)
	}
	return &Zap{
		l: logger,
	}
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
	m := map[Level]func(string, ...zap.Field){
		Debug: z.l.Debug,
		Info:  z.l.Info,
		Warn:  z.l.Warn,
		Error: z.l.Error,
		Fatal: z.l.Fatal,
		Panic: z.l.Panic,
	}
	return m[lvl]
}

func (Zap) convertFields(fs []Field) []zap.Field {
	zfs := make([]zap.Field, len(fs))
	for i, v := range fs {
		zfs[i] = v.zap
	}
	return zfs
}

func toZapLevel(lvl Level) zapcore.Level {
	m := map[Level]zapcore.Level{
		Debug: zapcore.DebugLevel,
		Info:  zapcore.InfoLevel,
		Warn:  zapcore.WarnLevel,
		Error: zapcore.ErrorLevel,
		Fatal: zapcore.FatalLevel,
		Panic: zapcore.PanicLevel,
	}
	return m[lvl]
}
