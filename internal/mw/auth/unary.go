package auth

import (
	"context"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor(parser tokenParser, scheme string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if shouldSkip(info.FullMethod) {
			return handler(ctx, req)
		}
		token, err := authFromCtx(ctx, scheme)
		if err != nil {
			return nil, err
		}
		userId, err := parser.GetUserId(ctx, auth.NewAccessToken(token))
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, core.UserIdKey, userId)
		return handler(ctx, req)
	}
}
