package biz

import (
	"context"
	"fmt"
	"reflect"

	"github.com/grimerssy/go-example/pkg/grpc_err"
)

type GreeterUseCase struct {
	users UserRepository
}

func NewGreeterUseCase(userRepository UserRepository) *GreeterUseCase {
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
