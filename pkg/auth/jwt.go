package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type ConfigJWT struct {
	AccessTokenMinutes time.Duration
	SigningString      string
}

type JWT struct {
	accessTokenTTL time.Duration
	signingString  string
}

func NewJWT(cfg ConfigJWT) *JWT {
	return &JWT{
		accessTokenTTL: cfg.AccessTokenMinutes * time.Minute,
		signingString:  cfg.SigningString,
	}
}

func (j *JWT) GenerateTokens(claims map[string]string) (Tokens, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.toClaims(claims))
	accessToken, err := token.SignedString([]byte(j.signingString))
	if err != nil {
		return nil, err
	}

	return NewTokens(accessToken), nil
}

func (j *JWT) ParseToken(token AccessToken) (map[string]string, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(j.signingString), nil
	}

	parsed, err := jwt.Parse(token.AccessToken(), keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt token: %w", err)
	}
	claims := parsed.Claims.(jwt.MapClaims)
	return j.toMap(claims), nil
}

func (j JWT) toClaims(m map[string]string) jwt.MapClaims {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(j.accessTokenTTL).Unix(),
		"iat": time.Now().Unix(),
	}
	for k, v := range m {
		claims[k] = v
	}
	return claims
}

func (JWT) toMap(claims jwt.MapClaims) map[string]string {
	m := make(map[string]string, len(claims))
	for k, v := range claims {
		m[k] = fmt.Sprintf("%v", v)
	}
	return m
}
