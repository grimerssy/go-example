package database

import (
	"github.com/google/wire"
)

var ProvideDB = wire.NewSet(
	NewPostgres,
)
