package biz

import (
	"context"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
)

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=biz -mock_names=TokenManager=tokenManagerMock,IdObfuscator=idObfuscatorMock,PasswordHasher=passwordHasherMock,UserRepository=userRepositoryMock
type TokenManager interface {
	GenerateToken(claims map[string]string) (auth.Token, error)
	ParseToken(token auth.AccessToken) (map[string]string, error)
}

type IdObfuscator interface {
	ObfuscateId(id int64) (int64, error)
	DeobfuscateId(obfuscated int64) (int64, error)
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	IsPasswordEqualToHash(password string, hash string) bool
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *core.User) error
	GetUserById(ctx context.Context, id int64) (*core.User, error)
	GetUserByName(ctx context.Context, name string) (*core.User, error)
	UpdateUserCount(ctx context.Context, user *core.User) error
}
