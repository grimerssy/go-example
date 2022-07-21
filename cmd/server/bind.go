package main

import (
	"github.com/google/wire"
	"github.com/grimerssy/go-example/configs"
	v1Api "github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grimerssy/go-example/internal/biz"
	"github.com/grimerssy/go-example/internal/data"
	"github.com/grimerssy/go-example/internal/mw"
	mwAuth "github.com/grimerssy/go-example/internal/mw/auth"
	mwLog "github.com/grimerssy/go-example/internal/mw/log"
	"github.com/grimerssy/go-example/internal/server"
	"github.com/grimerssy/go-example/internal/service"
	"github.com/grimerssy/go-example/internal/service/v1"
	"github.com/grimerssy/go-example/pkg/auth"
	"github.com/grimerssy/go-example/pkg/database"
	"github.com/grimerssy/go-example/pkg/hash"
	"github.com/grimerssy/go-example/pkg/id"
	"github.com/grimerssy/go-example/pkg/log"
)

var withBind = wire.NewSet(
	configs.ProvideConfig,
	server.ProvideServer,
	mw.ProvideInterceptors,
	service.ProvideServices,
	wire.Bind(new(v1Api.AuthServiceServer), new(*v1.AuthService)),
	wire.Bind(new(v1Api.GreeterServiceServer), new(*v1.GreeterService)),
	biz.ProvideUseCases,
	wire.Bind(new(v1.AuthUseCase), new(*biz.AuthUseCase)),
	wire.Bind(new(mwAuth.TokenParser), new(*biz.AuthUseCase)),
	wire.Bind(new(v1.GreeterUseCase), new(*biz.GreeterUseCase)),
	data.ProvideRepositories,
	wire.Bind(new(biz.UserRepository), new(*data.UserRepository)),
	auth.ProvideAuth,
	wire.Bind(new(biz.TokenManager), new(*auth.JWT)),
	database.ProvideDB,
	hash.ProvideHash,
	wire.Bind(new(biz.PasswordHasher), new(*hash.Bcrypt)),
	id.ProvideId,
	wire.Bind(new(biz.IdObfuscator), new(*id.Optimus)),
	log.ProvideLog,
	wire.Bind(new(mwLog.Logger), new(*log.Zap)),
)
