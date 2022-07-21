// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/grimerssy/go-example/configs"
	"github.com/grimerssy/go-example/internal/biz"
	"github.com/grimerssy/go-example/internal/data"
	"github.com/grimerssy/go-example/internal/mw"
	"github.com/grimerssy/go-example/internal/server"
	"github.com/grimerssy/go-example/internal/service/v1"
	"github.com/grimerssy/go-example/pkg/auth"
	"github.com/grimerssy/go-example/pkg/database"
	"github.com/grimerssy/go-example/pkg/hash"
	"github.com/grimerssy/go-example/pkg/id"
	"github.com/grimerssy/go-example/pkg/log"
)

// Injectors from wire.go:

func InitializeServer(file string) *server.Server {
	config := configs.NewConfig(file)
	configServer := config.Server
	configMW := config.MW
	configJWT := config.JWT
	jwt := auth.NewJWT(configJWT)
	configOptimus := config.Optimus
	optimus := id.NewOptimus(configOptimus)
	configBcrypt := config.Bcrypt
	bcrypt := hash.NewBcrypt(configBcrypt)
	configDB := config.DB
	db := database.NewPostgres(configDB)
	userRepository := data.NewUserRepository(db)
	authUseCase := biz.NewAuthUseCase(jwt, optimus, bcrypt, userRepository)
	configZap := config.Zap
	zap := log.NewZap(configZap)
	v := mw.NewUnaryServerInterceptors(configMW, authUseCase, zap)
	v2 := mw.NewStreamServerInterceptors()
	v3 := server.NewGrpcServerOptions(v, v2)
	authService := v1.NewAuthService(authUseCase)
	greeterUseCase := biz.NewGreeterUseCase(userRepository)
	greeterService := v1.NewGreeterService(greeterUseCase)
	v4 := server.NewRegisterServicesFunc(authService, greeterService)
	grpcServer := server.NewGrpcServer(v3, v4)
	handler := server.NewHttpHandler(configServer)
	serverServer := server.NewServer(configServer, grpcServer, handler)
	return serverServer
}
