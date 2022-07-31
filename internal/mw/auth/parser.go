package auth

import (
	"context"

	"github.com/grimerssy/go-example/pkg/auth"
)

//go:generate mockgen -source=parser.go -destination=parser_mock.go -package=auth -mock_names=TokenParser=tokenParserMock
type TokenParser interface {
	GetUserId(ctx context.Context, tokens auth.AccessToken) (int64, error)
}
