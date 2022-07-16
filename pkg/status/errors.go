package status

import (
	"context"
	"errors"

	"github.com/grimerssy/go-example/pkg/consts"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapError(err error) error {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return newStatusError(codes.DeadlineExceeded,
			context.DeadlineExceeded.Error(), err.Error())

	case errors.Is(err, context.Canceled):
		return newStatusError(codes.Canceled,
			context.Canceled.Error(), err.Error())

	case errors.Is(err, consts.ErrInvalidPassword):
		return newStatusError(codes.Unauthenticated,
			consts.ErrInvalidPassword.Error(), err.Error())

	case errors.Is(err, consts.ErrUserAlreadyExists):
		return newStatusError(codes.AlreadyExists,
			consts.ErrUserAlreadyExists.Error(), err.Error())

	case errors.Is(err, consts.ErrUserNotFound):
		return newStatusError(codes.NotFound,
			consts.ErrUserNotFound.Error(), err.Error())

	case errors.Is(err, consts.ErrContextHasNoUserId):
		return newStatusError(codes.Internal,
			consts.ErrContextHasNoUserId.Error(), err.Error())

	default:
		return newStatusError(codes.Unknown,
			consts.UnexpectedErrorMessage, err.Error())
	}
}

type statusError struct {
	s      *status.Status
	logMsg string
}

func newStatusError(code codes.Code, errMsg, logMsg string) *statusError {
	return &statusError{
		s:      status.New(code, errMsg),
		logMsg: logMsg,
	}
}

func (e *statusError) Error() string {
	return e.s.Err().Error()
}

func (e *statusError) GRPCStatus() *status.Status {
	return e.s
}

func (e *statusError) Is(target error) bool {
	return errors.Is(e.s.Err(), target)
}

func (e *statusError) LogMessage() string {
	return e.logMsg
}
