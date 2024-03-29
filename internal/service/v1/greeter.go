package v1

import (
	"context"
	"reflect"

	v1 "github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/grpc_err"
)

//go:generate mockgen -source=greeter.go -destination=greeter_mock.go -package=v1 -mock_names=GreeterUseCase=greeterUseCaseMock
type GreeterUseCase interface {
	Greet(ctx context.Context, userId int64) (string, error)
}

type GreeterService struct {
	uc GreeterUseCase
}

func NewGreeterService(greeterUseCase GreeterUseCase) *GreeterService {
	if reflect.ValueOf(greeterUseCase).IsNil() {
		panic("greeterUseCase cannot be nil")
	}
	return &GreeterService{
		uc: greeterUseCase,
	}
}

func (s *GreeterService) Greet(ctx context.Context, req *v1.GreetRequest,
) (*v1.GreetResponse, error) {
	userId, ok := ctx.Value(core.UserIdKey).(int64)
	if !ok {
		return nil, grpc_err.ContextHasNoValue("user id", 0)
	}
	message, err := s.uc.Greet(ctx, userId)
	if err != nil {
		return nil, grpc_err.Wrap(err, 0)
	}
	res := &v1.GreetResponse{
		Message: message,
	}
	return res, nil
}
