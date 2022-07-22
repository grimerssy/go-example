package data

import (
	"testing"

	"github.com/grimerssy/go-example/pkg/grpc_err"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data access suite")
}

var (
	errUserAlreadyExists = grpc_err.AlreadyExists("user", 0)
	errUserNotFound      = grpc_err.NotFound("user", 0)
)
