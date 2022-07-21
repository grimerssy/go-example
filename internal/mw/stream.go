package mw

import (
	"google.golang.org/grpc"
)

func NewStreamServerInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{}
}
