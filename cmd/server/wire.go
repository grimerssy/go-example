//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/grimerssy/go-example/internal/server"
)

func InitializeServer(file string) *server.Server {
	panic(wire.Build(withBind))
}
