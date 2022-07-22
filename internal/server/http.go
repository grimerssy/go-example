package server

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/grimerssy/go-example/internal/api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewHttpHandler(cfg ConfigServer, dopts []grpc.DialOption) http.Handler {
	mux := http.NewServeMux()
	ep := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	gw := getGatewayHandler(ep, dopts)

	mux.Handle("/", gw)
	return mux
}

func getGatewayHandler(endpoint string, dopts []grpc.DialOption) http.Handler {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	err := v1.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, endpoint, dopts)
	if err != nil {
		panic(err)
	}
	err = v1.RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, endpoint, dopts)
	if err != nil {
		panic(err)
	}
	return mux
}
