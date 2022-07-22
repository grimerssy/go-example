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
	"github.com/spf13/viper"
)

const (
	configDir = "configs/"
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
	cfg := new(Config)
	files := getConfigFiles(env)
	if err := loadConfig(cfg, files); err != nil {
		panic(err)
	}
	return cfg
}

func loadConfig(cfg *Config, files []string) error {
	for _, f := range files {
		viper.SetConfigFile(configDir + f)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		if err := viper.Unmarshal(cfg); err != nil {
			return err
		}
	}
	return nil
}

func getConfigFiles(env core.Environment) []string {
	m := map[core.Environment][]string{
		core.Development: {"dev.yaml", "dev.env"},
		core.Staging:     {"stage.yaml", "stage.env"},
		core.Production:  {"prod.yaml", "prod.env"},
	}
	return m[env]
}
