//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/grimerssy/go-example/internal/core"
	"github.com/grimerssy/go-example/internal/server"

	// used by wire
	// if not imported, go mod tidy deletes go.sum entry
	_ "github.com/google/subcommands"
	_ "golang.org/x/tools/go/ast/astutil"
	_ "golang.org/x/tools/go/packages"
	_ "golang.org/x/tools/go/types/typeutil"
)

func initializeServer(env core.Environment) (*server.Server, func()) {
	panic(wire.Build(withBind))
}
