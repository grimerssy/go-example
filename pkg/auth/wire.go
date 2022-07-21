package auth

import (
	"github.com/google/wire"
)

var ProvideAuth = wire.NewSet(
	NewJWT,
)
