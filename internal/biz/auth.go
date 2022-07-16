package biz

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
	"github.com/grimerssy/go-example/pkg/consts"
)

type AuthUseCase struct {
	tokens    tokenManager
	ids       idObfuscator
	passwords passwordHasher
	users     userRepository
}

func NewAuthUseCase(
	tokenManager tokenManager,
	idObfuscator idObfuscator,
	passwordHasher passwordHasher,
	userRepository userRepository) *AuthUseCase {
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
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword
	err = uc.users.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (uc *AuthUseCase) Login(ctx context.Context, input *core.User) (auth.Tokens, error) {
	user, err := uc.users.GetUserByName(ctx, input.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if !uc.passwords.IsPasswordEqualToHash(input.Password, user.Password) {
		return nil, consts.ErrInvalidPassword
	}
	obfuscatedId, err := uc.ids.ObfuscateId(user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to obfuscate user id: %w", err)
	}
	claims := newUserIdClaims(uc.tokens.DefaultClaims(), obfuscatedId)
	tokens, err := uc.tokens.GenerateTokens(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}
	return tokens, nil
}

func (uc *AuthUseCase) GetUserId(ctx context.Context, tokens auth.Tokens) (int64, error) {
	claims, err := uc.tokens.ParseTokens(tokens, &userIdClaims{})
	if err != nil {
		return 0, fmt.Errorf("failed to parse tokens: %w", err)
	}
	userIdClaims, ok := claims.(*userIdClaims)
	if !ok {
		return 0, errors.New("failed to cast claims to type userIdClaims")
	}
	obfuscatedId := userIdClaims.UserId
	userId, err := uc.ids.DeobfuscateId(obfuscatedId)
	if err != nil {
		return 0, fmt.Errorf("failed to deobfuscate user id: %w", err)
	}
	return userId, nil
}
