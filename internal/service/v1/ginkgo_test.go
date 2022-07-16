package v1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/status"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Presentation suite")
}

func statusFromError(err error) *status.Status {
	st, ok := status.FromError(err)
	Expect(ok).To(BeTrue())
	return st
}
