package auth

import (
	"context"

	"github.com/grimerssy/go-example/pkg/auth"
)

//go:generate mockery --name=TokenParser --with-expecter --quiet
type TokenParser interface {
	GetUserId(ctx context.Context, tokens auth.AccessToken) (int64, error)
}
