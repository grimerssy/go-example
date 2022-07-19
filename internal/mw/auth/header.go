package auth

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
)

func authFromCtx(ctx context.Context, scheme string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Internal, "failed to extract metadata")
	}
	vs := md.Get(authorizationHeader)
	if len(vs) == 0 {
		return "", badAuthHeader("no header was provided")
	}
	split := strings.SplitN(vs[0], " ", 2)
	if len(split) != 2 {
		return "", badAuthHeader("invalid authorization string")
	}
	if !strings.EqualFold(split[0], scheme) {
		return "", badAuthHeader(
			fmt.Sprintf("expected token with %s scheme", scheme))
	}
	return split[1], nil
}

func badAuthHeader(cause string) error {
	return status.Error(codes.Unauthenticated,
		fmt.Sprintf("bad Authorization header: %s", cause))
}
