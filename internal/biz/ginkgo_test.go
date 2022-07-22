package biz

import (
	"testing"

	"github.com/grimerssy/go-example/pkg/grpc_err"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_UseCases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Business logic suite")
}

var (
	errUserAlreadyExists = grpc_err.AlreadyExists("user", 0)
	errUserNotFound      = grpc_err.NotFound("user", 0)
	errInvalidPassword   = grpc_err.InvalidPassword(0)
)
