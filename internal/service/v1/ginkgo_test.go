package v1

import (
	"testing"

	"github.com/grimerssy/go-example/pkg/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/status"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Presentation suite")
}

var (
	errUserAlreadyExists  = errors.AlreadyExists("user", 0)
	errUserNotFound       = errors.NotFound("user", 0)
	errInvalidPassword    = errors.InvalidPassword(0)
	errContextHasNoUserId = errors.ContextHasNoValue("user id", 0)
)

func statusFromError(err error) *status.Status {
	st, ok := status.FromError(err)
	Expect(ok).To(BeTrue())
	return st
}
