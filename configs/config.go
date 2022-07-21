package configs

import (
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/internal/mw"
	"github.com/grimerssy/go-example/internal/server"
	"github.com/grimerssy/go-example/pkg/auth"
	"github.com/grimerssy/go-example/pkg/database"
	"github.com/grimerssy/go-example/pkg/hash"
	"github.com/grimerssy/go-example/pkg/id"
	"github.com/grimerssy/go-example/pkg/log"
)

type Config struct {
	Server  server.ConfigServer
	MW      mw.ConfigMW
	JWT     auth.ConfigJWT
	DB      database.ConfigDB
	Bcrypt  hash.ConfigBcrypt
	Optimus id.ConfigOptimus
	Zap     log.ConfigZap
}

func NewConfig(env core.Environment) *Config {
	return &Config{}
}
