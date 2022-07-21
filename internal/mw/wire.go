package mw

import (
	"github.com/google/wire"
)

var ProvideInterceptors = wire.NewSet(
	NewUnaryServerInterceptors,
	NewStreamServerInterceptors,
)
