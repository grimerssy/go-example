package biz

import (
	"github.com/grimerssy/go-example/pkg/auth"
)

type userIdClaims struct {
	auth.Claims
	UserId int64
}

func newUserIdClaims(claims auth.Claims, userId int64) *userIdClaims {
	return &userIdClaims{
		Claims: claims,
		UserId: userId,
	}
}
