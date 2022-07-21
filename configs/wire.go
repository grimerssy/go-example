package configs

import (
	"github.com/google/wire"
)

var ProvideConfig = wire.NewSet(
	NewConfig,
	wire.FieldsOf(new(*Config),
		"Server",
		"MW",
		"JWT",
		"DB",
		"Bcrypt",
		"Optimus",
		"Zap",
	),
)
