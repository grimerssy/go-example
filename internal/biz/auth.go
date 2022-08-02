package biz

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
	"github.com/grimerssy/go-example/pkg/grpc_err"
)

//go:generate mockgen -source=auth.go -destination=auth_mock.go -package=biz -mock_names=TokenManager=tokenManagerMock,IdObfuscator=idObfuscatorMock,PasswordHasher=passwordHasherMock,AuthUserRepository=authUserRepositoryMock
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

type AuthUserRepository interface {
	CreateUser(ctx context.Context, user *core.User) error
	GetUserByName(ctx context.Context, name string) (*core.User, error)
}

type AuthUseCase struct {
	tokens    TokenManager
	ids       IdObfuscator
	passwords PasswordHasher
	users     AuthUserRepository
}

func NewAuthUseCase(
	tokenManager TokenManager,
	idObfuscator IdObfuscator,
	passwordHasher PasswordHasher,
	userRepository AuthUserRepository,
) *AuthUseCase {
	switch {
	case reflect.ValueOf(idObfuscator).IsNil():
		panic("idObfuscator cannot be nil")
	case reflect.ValueOf(tokenManager).IsNil():
		panic("tokenManager cannot be nil")
	case reflect.ValueOf(passwordHasher).IsNil():
		panic("passwordHasher cannot be nil")
	case reflect.ValueOf(userRepository).IsNil():
		panic("userRepository cannot be nil")
	}
	return &AuthUseCase{
		ids:       idObfuscator,
		tokens:    tokenManager,
		passwords: passwordHasher,
		users:     userRepository,
	}
}

func (uc *AuthUseCase) Signup(ctx context.Context, user *core.User) error {
	hashedPassword, err := uc.passwords.HashPassword(user.Password)
	if err != nil {
		return grpc_err.Wrap(err, 0)
	}
	user.Password = hashedPassword
	err = uc.users.CreateUser(ctx, user)
	if err != nil {
		return grpc_err.Wrap(err, 0)
	}
	return nil
}

func (uc *AuthUseCase) Login(ctx context.Context, input *core.User,
) (auth.Token, error) {
	user, err := uc.users.GetUserByName(ctx, input.Name)
	if err != nil {
		return nil, grpc_err.Wrap(err, 0)
	}
	if !uc.passwords.IsPasswordEqualToHash(input.Password, user.Password) {
		return nil, grpc_err.InvalidPassword(0)
	}
	obfuscatedId, err := uc.ids.ObfuscateId(user.Id)
	if err != nil {
		return nil, grpc_err.Wrap(err, 0)
	}
	claims := map[string]string{
		core.UserIdKey: fmt.Sprintf("%v", obfuscatedId),
	}
	token, err := uc.tokens.GenerateToken(claims)
	if err != nil {
		return nil, grpc_err.Wrap(err, 0)
	}
	return token, nil
}

func (uc *AuthUseCase) GetUserId(ctx context.Context, token auth.AccessToken,
) (int64, error) {
	claims, err := uc.tokens.ParseToken(token)
	if err != nil {
		return 0, grpc_err.Wrap(err, 0)
	}
	obfuscatedId, err := strconv.ParseInt(claims[core.UserIdKey], 10, 64)
	if err != nil {
		return 0, grpc_err.Wrap(err, 0)
	}
	userId, err := uc.ids.DeobfuscateId(obfuscatedId)
	if err != nil {
		return 0, grpc_err.Wrap(err, 0)
	}
	return userId, nil
}
