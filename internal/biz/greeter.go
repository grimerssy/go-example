package biz

import (
	"context"
	"fmt"
	"reflect"
)

type GreeterUseCase struct {
	users userRepository
}

func NewGreeterUseCase(userRepository userRepository) *GreeterUseCase {
	if reflect.ValueOf(userRepository).IsNil() {
		panic("userRepository cannot be nil")
	}
	return &GreeterUseCase{
		users: userRepository,
	}
}

func (uc *GreeterUseCase) Greet(ctx context.Context, userId int64) (string, error) {
	user, err := uc.users.GetUserById(ctx, userId)
	if err != nil {
		return "", fmt.Errorf("failed to get user by id: %w", err)
	}
	user.Count++
	err = uc.users.UpdateUserCount(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to update user count: %w", err)
	}
	return fmt.Sprintf("Glad to see you! It's your %d time here", user.Count), nil
}
