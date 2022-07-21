//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/internal/server"
)

func initializeServer(env core.Environment) (*server.Server, func()) {
	panic(wire.Build(withBind))
}
