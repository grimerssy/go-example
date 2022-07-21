package id

import (
	"github.com/google/wire"
)

var ProvideId = wire.NewSet(
	NewOptimus,
)
