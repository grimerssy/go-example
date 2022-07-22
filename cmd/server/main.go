package main

import (
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grimerssy/go-example/internal/core"
	_ "github.com/lib/pq"
)

func main() {
	env := getEnvironment()
	server, cleanup := initializeServer(env)
	defer cleanup()

	go func() {
		err := server.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(); err != nil {
		panic(err)
	}
}

func getEnvironment() core.Environment {
	isDev := flag.Bool("dev", false, "set development environment")
	isStage := flag.Bool("stage", false, "set staging environment")
	isProd := flag.Bool("prod", false, "set production environment")
	flag.Parse()

	switch {
	case *isDev:
		return core.Development
	case *isStage:
		return core.Staging
	case *isProd:
		return core.Production
	default:
		panic("exactly one environment flag must be set (-dev/-stage/-prod)")
	}
}
