package mw

import (
	"github.com/grimerssy/go-example/internal/mw/auth"
	"github.com/grimerssy/go-example/internal/mw/log"
	"github.com/grimerssy/go-example/internal/mw/validation"
	"google.golang.org/grpc"
)

func NewUnaryServerInterceptors(tp auth.TokenParser, l log.Logger, cfg ConfigMW,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		validation.UnaryServerInterceptor(),
		auth.UnaryServerInterceptor(tp, cfg.AuthScheme),
		log.UnaryServerInterceptor(l),
	}
}
