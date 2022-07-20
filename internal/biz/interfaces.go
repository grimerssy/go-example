package biz

import (
	"context"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
)

//go:generate mockery --name=TokenManager --with-expecter
type TokenManager interface {
	DefaultClaims() auth.Claims
	GenerateTokens(claims auth.Claims) (auth.Tokens, error)
	ParseToken(token auth.AccessToken, claims auth.Claims) (auth.Claims, error)
}

//go:generate mockery --name=IdObfuscator --with-expecter
type IdObfuscator interface {
	ObfuscateId(id int64) (int64, error)
	DeobfuscateId(obfuscated int64) (int64, error)
}

//go:generate mockery --name=PasswordHasher --with-expecter
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	IsPasswordEqualToHash(password string, hash string) bool
}

//go:generate mockery --name=UserRepository --with-expecter
type UserRepository interface {
	CreateUser(ctx context.Context, user *core.User) error
	GetUserById(ctx context.Context, id int64) (*core.User, error)
	GetUserByName(ctx context.Context, name string) (*core.User, error)
	UpdateUserCount(ctx context.Context, user *core.User) error
}
