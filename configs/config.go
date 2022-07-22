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
	fileName := getFileName(env)
	if err := loadConfigFile(cfg, fileName); err != nil {
		panic(err)
	}
	envPrefix := getEnvPrefix(env)
	if err := loadEnvVars(cfg, envPrefix); err != nil {
		panic(err)
	}
	return cfg
}

func loadConfigFile(cfg *Config, fileName string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName(fileName)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(cfg)
}

func loadEnvVars(cfg *Config, envPrefix string) error {
	viper.SetEnvPrefix(envPrefix)
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(cfg)
}

func getFileName(env core.Environment) string {
	m := map[core.Environment]string{
		core.Development: "dev",
		core.Staging:     "stage",
		core.Production:  "prod",
	}
	return m[env]
}

func getEnvPrefix(env core.Environment) string {
	m := map[core.Environment]string{
		core.Development: "DEV",
		core.Staging:     "STAGE",
		core.Production:  "PROD",
	}
	return m[env]
}
