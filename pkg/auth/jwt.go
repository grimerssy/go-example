package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type ConfigJWT struct {
	TokenMinutes  time.Duration
	SigningString string
}

type JWT struct {
	tokenTTL      time.Duration
	signingString string
}

func NewJWT(cfg ConfigJWT) *JWT {
	return &JWT{
		tokenTTL:      cfg.TokenMinutes * time.Minute,
		signingString: cfg.SigningString,
	}
}

func (a *JWT) DefaultClaims() Claims {
	return jwt.StandardClaims{
		ExpiresAt: time.Now().Add(a.tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}
}

func (a *JWT) GenerateTokens(claims Claims) (Tokens, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(a.signingString))
	if err != nil {
		return nil, err
	}

	return newTokens(accessToken), nil
}

func (a *JWT) ParseTokens(tokens Tokens, claims Claims) (Claims, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(a.signingString), nil
	}

	token, err := jwt.ParseWithClaims(tokens.AccessToken(), claims, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt token: %w", err)
	}

	return token.Claims, nil
}
