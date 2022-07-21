package service

import (
	"github.com/google/wire"
	v1 "github.com/grimerssy/go-example/internal/service/v1"
)

var ProvideServices = wire.NewSet(
	v1.NewAuthService,
	v1.NewGreeterService,
)
