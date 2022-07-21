package log

import (
	"github.com/google/wire"
)

var ProvideLog = wire.NewSet(
	NewZap,
)
