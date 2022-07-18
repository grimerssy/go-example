package v1

import (
	"context"
	"reflect"

	"github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/auth"
	"github.com/grimerssy/go-example/pkg/errors"
)

//go:generate mockery --name=authUseCase --with-expecter --exported
type authUseCase interface {
	Signup(ctx context.Context, user *core.User) error
	Login(ctx context.Context, input *core.User) (auth.Tokens, error)
}

type AuthService struct {
	uc authUseCase
}

func NewAuthService(authUseCase authUseCase) *AuthService {
	if reflect.ValueOf(authUseCase).IsNil() {
		panic("authUseCase cannot be nil")
	}
	return &AuthService{
		uc: authUseCase,
	}
}

func (s *AuthService) Signup(ctx context.Context,
	req *v1.SignupRequest) (*v1.SignupResponse, error) {
	user := &core.User{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}
	err := s.uc.Signup(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	return &v1.SignupResponse{}, nil
}

func (s *AuthService) Login(ctx context.Context,
	req *v1.LoginRequest) (*v1.LoginResponse, error) {
	user := &core.User{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}
	tokens, err := s.uc.Login(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}
	res := &v1.LoginResponse{
		AccessToken: tokens.AccessToken(),
	}
	return res, nil
}
