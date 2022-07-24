package mw

import (
	"github.com/grimerssy/go-example/internal/mw/auth"
	"github.com/grimerssy/go-example/internal/mw/log"
	"github.com/grimerssy/go-example/internal/mw/validation"
	"google.golang.org/grpc"
)

func NewUnaryServerInterceptors(cfg ConfigMW, tp auth.TokenParser, l log.Logger,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		log.UnaryServerInterceptor(l),
		validation.UnaryServerInterceptor(),
		auth.UnaryServerInterceptor(tp, cfg.AuthScheme),
	}
}
