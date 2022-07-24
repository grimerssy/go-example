package biz

import (
	"context"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
)

//go:generate mockery --name=TokenManager --with-expecter --quiet
type TokenManager interface {
	GenerateTokens(claims map[string]string) (auth.Tokens, error)
	ParseToken(token auth.AccessToken) (map[string]string, error)
}

//go:generate mockery --name=IdObfuscator --with-expecter --quiet
type IdObfuscator interface {
	ObfuscateId(id int64) (int64, error)
	DeobfuscateId(obfuscated int64) (int64, error)
}

//go:generate mockery --name=PasswordHasher --with-expecter --quiet
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	IsPasswordEqualToHash(password string, hash string) bool
}

//go:generate mockery --name=UserRepository --with-expecter --quiet
type UserRepository interface {
	CreateUser(ctx context.Context, user *core.User) error
	GetUserById(ctx context.Context, id int64) (*core.User, error)
	GetUserByName(ctx context.Context, name string) (*core.User, error)
	UpdateUserCount(ctx context.Context, user *core.User) error
}
