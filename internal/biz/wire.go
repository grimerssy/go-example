package biz

import (
	"github.com/google/wire"
)

var ProvideUseCases = wire.NewSet(
	NewAuthUseCase,
	NewGreeterUseCase,
)
