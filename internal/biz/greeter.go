package biz

import (
	"context"
	"fmt"
	"reflect"

	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/pkg/grpc_err"
)

//go:generate mockgen -source=greeter.go -destination=greeter_mock.go -package=biz -mock_names=GreeterUserRepository=greeterUserRepositoryMock
type GreeterUserRepository interface {
	GetUserById(ctx context.Context, id int64) (*core.User, error)
	UpdateUserCount(ctx context.Context, user *core.User) error
}

type GreeterUseCase struct {
	users GreeterUserRepository
}

func NewGreeterUseCase(userRepository GreeterUserRepository) *GreeterUseCase {
	if reflect.ValueOf(userRepository).IsNil() {
		panic("userRepository cannot be nil")
	}
	return &GreeterUseCase{
		users: userRepository,
	}
}

func (uc *GreeterUseCase) Greet(ctx context.Context, userId int64,
) (string, error) {
	user, err := uc.users.GetUserById(ctx, userId)
	if err != nil {
		return "", grpc_err.Wrap(err, 0)
	}
	user.Count++
	err = uc.users.UpdateUserCount(ctx, user)
	if err != nil {
		return "", grpc_err.Wrap(err, 0)
	}
	return fmt.Sprintf("Glad to see you, %s! It's your %d time here",
		user.Name, user.Count), nil
}
