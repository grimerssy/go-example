package v1

import (
	"testing"

	"github.com/grimerssy/go-example/pkg/grpc_err"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/status"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Presentation suite")
}

var (
	errUserAlreadyExists  = grpc_err.AlreadyExists("user", 0)
	errUserNotFound       = grpc_err.NotFound("user", 0)
	errInvalidPassword    = grpc_err.InvalidPassword(0)
	errContextHasNoUserId = grpc_err.ContextHasNoValue("user id", 0)
)

func statusFromError(err error) *status.Status {
	st, ok := status.FromError(err)
	Expect(ok).To(BeTrue())
	return st
}
