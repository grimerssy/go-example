package auth

import (
	"context"

	"github.com/grimerssy/go-example/pkg/auth"
)

type tokenParser interface {
	GetUserId(ctx context.Context, tokens auth.AccessToken) (int64, error)
}
