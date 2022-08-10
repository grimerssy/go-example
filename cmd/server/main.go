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
		if !errors.Is(err, http.ErrServerClosed) {
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

	setCount := 0
	for _, flag := range []*bool{isDev, isStage, isProd} {
		if *flag {
			setCount++
		}
	}

	switch {
	case setCount != 1:
		panic("exactly one environment flag must be set (-dev/-stage/-prod)")
	case *isDev:
		return core.Development
	case *isStage:
		return core.Staging
	case *isProd:
		return core.Production
	default:
		panic("did not match any of the env flags")
	}
}
