package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type ConfigServer struct {
	ShutdownSeconds time.Duration
	Host            string
	Port            int
}

type Server struct {
	shutdownTO time.Duration
	http       *http.Server
}

func NewServer(cfg ConfigServer, srv *grpc.Server, handler http.Handler,
) *Server {
	return &Server{
		shutdownTO: cfg.ShutdownSeconds * time.Second,
		http: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: handleGrpc(srv, handler),
		},
	}
}

func (s *Server) Run() error {
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTO)
	defer cancel()
	return s.http.Shutdown(ctx)
}

func handleGrpc(srv *grpc.Server, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 &&
			strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			srv.ServeHTTP(w, r)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}
