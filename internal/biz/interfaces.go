package biz

import (
	"context"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
)

//go:generate mockery --name=tokenManager --with-expecter --exported
type tokenManager interface {
	DefaultClaims() auth.Claims
	GenerateTokens(claims auth.Claims) (auth.Tokens, error)
	ParseTokens(tokens auth.Tokens, claims auth.Claims) (auth.Claims, error)
}

//go:generate mockery --name=idObfuscator --with-expecter --exported
type idObfuscator interface {
	ObfuscateId(id int64) (int64, error)
	DeobfuscateId(obfuscated int64) (int64, error)
}

//go:generate mockery --name=passwordHasher --with-expecter --exported
type passwordHasher interface {
	HashPassword(password string) (string, error)
	IsPasswordEqualToHash(password string, hash string) bool
}

//go:generate mockery --name=userRepository --with-expecter --exported
type userRepository interface {
	CreateUser(ctx context.Context, user *core.User) error
	GetUserById(ctx context.Context, id int64) (*core.User, error)
	GetUserByName(ctx context.Context, name string) (*core.User, error)
	UpdateUserCount(ctx context.Context, user *core.User) error
}
