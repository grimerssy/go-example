package server

import (
	"github.com/grimerssy/go-example/internal/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcServer(opts []grpc.ServerOption, regSrv func(srv *grpc.Server),
) *grpc.Server {
	srv := grpc.NewServer(opts...)
	regSrv(srv)
	return srv
}

func NewGrpcServerOptions(unaryInts []grpc.UnaryServerInterceptor,
	streamInts []grpc.StreamServerInterceptor,
) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(unaryInts...),
		grpc.ChainStreamInterceptor(streamInts...),
	}
}

func NewGrpcDialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
}

func NewRegisterServicesFunc(
	v1Auth v1.AuthServiceServer,
	v1Greeter v1.GreeterServiceServer,
) func(srv *grpc.Server) {
	return func(srv *grpc.Server) {
		v1.RegisterAuthServiceServer(srv, v1Auth)
		v1.RegisterGreeterServiceServer(srv, v1Greeter)
	}
}
