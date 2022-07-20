package biz

import (
	"context"
	"reflect"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
	"github.com/grimerssy/go-example/pkg/errors"
)

type AuthUseCase struct {
	tokens    TokenManager
	ids       IdObfuscator
	passwords PasswordHasher
	users     UserRepository
}

func NewAuthUseCase(
	tokenManager TokenManager,
	idObfuscator IdObfuscator,
	passwordHasher PasswordHasher,
	userRepository UserRepository,
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
		return errors.Wrap(err, 0)
	}
	user.Password = hashedPassword
	err = uc.users.CreateUser(ctx, user)
	if err != nil {
		return errors.Wrap(err, 0)
	}
	return nil
}

func (uc *AuthUseCase) Login(ctx context.Context, input *core.User,
) (auth.Tokens, error) {
	user, err := uc.users.GetUserByName(ctx, input.Name)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	if !uc.passwords.IsPasswordEqualToHash(input.Password, user.Password) {
		return nil, errors.InvalidPassword(0)
	}
	obfuscatedId, err := uc.ids.ObfuscateId(user.Id)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	claims := newUserIdClaims(uc.tokens.DefaultClaims(), obfuscatedId)
	tokens, err := uc.tokens.GenerateTokens(claims)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	return tokens, nil
}

func (uc *AuthUseCase) GetUserId(ctx context.Context, token auth.AccessToken,
) (int64, error) {
	claims, err := uc.tokens.ParseToken(token, &userIdClaims{})
	if err != nil {
		return 0, errors.Wrap(err, 0)
	}
	userIdClaims := claims.(*userIdClaims)
	obfuscatedId := userIdClaims.UserId
	userId, err := uc.ids.DeobfuscateId(obfuscatedId)
	if err != nil {
		return 0, errors.Wrap(err, 0)
	}
	return userId, nil
}
