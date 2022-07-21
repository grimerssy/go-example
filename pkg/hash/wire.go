package hash

import (
	"github.com/google/wire"
)

var ProvideHash = wire.NewSet(
	NewBcrypt,
)
