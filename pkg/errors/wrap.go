package errors

import (
	"runtime"

	"github.com/grimerssy/go-example/pkg/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Wrap(err error, offset int) error {
	caller := getCaller(offset + 1)
	if w, ok := err.(*wrappedError); ok {
		w.callers = append(w.callers, caller)
		return w
	}
	return &wrappedError{
		code:    status.Code(err),
		msg:     err.Error(),
		callers: []string{caller},
	}
}

type wrappedError struct {
	code    codes.Code
	msg     string
	callers []string
}

func (e *wrappedError) Error() string {
	return e.msg
}

func (e *wrappedError) Is(target error) bool {
	if w, ok := target.(*wrappedError); ok {
		return e.code == w.code && e.msg == w.msg
	}
	return e.code == codes.Unknown && e.msg == target.Error()
}

func (e *wrappedError) GRPCStatus() *status.Status {
	return status.New(e.code, e.msg)
}

func (e *wrappedError) Callers() []string {
	return slices.ReverseCopy(e.callers)
}

func getCaller(n int) string {
	pc, _, _, ok := runtime.Caller(n + 1)
	if !ok {
		panic("failed to get caller")
	}
	return runtime.FuncForPC(pc).Name()
}
