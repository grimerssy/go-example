package grpc_err

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

func AlreadyExists(what string, offset int) error {
	return &wrappedError{
		code:    codes.AlreadyExists,
		msg:     fmt.Sprintf("%s already exists", what),
		callers: []string{getCaller(offset + 1)},
	}
}

func NotFound(what string, offset int) error {
	return &wrappedError{
		code:    codes.NotFound,
		msg:     fmt.Sprintf("%s was not found", what),
		callers: []string{getCaller(offset + 1)},
	}
}

func InvalidPassword(offset int) error {
	return &wrappedError{
		code:    codes.Unauthenticated,
		msg:     "invalid password",
		callers: []string{getCaller(offset + 1)},
	}
}

func ContextHasNoValue(what string, offset int) error {
	return &wrappedError{
		code:    codes.Internal,
		msg:     fmt.Sprintf("context has no %s", what),
		callers: []string{getCaller(offset + 1)},
	}
}
