package server

import (
	"github.com/google/wire"
)

var ProvideServer = wire.NewSet(
	NewServer,
	NewGrpcServer,
	NewGrpcServerOptions,
	NewRegisterServicesFunc,
	NewHttpHandler,
	NewGrpcDialOptions,
)
