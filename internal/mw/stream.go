package mw

import (
	"google.golang.org/grpc"
)

func NewStreamServerInterceptors() []grpc.StreamServerInterceptor {
	return []grpc.StreamServerInterceptor{}
}
