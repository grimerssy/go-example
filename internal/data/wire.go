package data

import (
	"github.com/google/wire"
)

var ProvideRepositories = wire.NewSet(
	NewUserRepository,
)
