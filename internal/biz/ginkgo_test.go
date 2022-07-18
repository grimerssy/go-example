package biz

import (
	"testing"

	"github.com/grimerssy/go-example/pkg/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_UseCases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Business logic suite")
}

var (
	errUserAlreadyExists = errors.AlreadyExists("user", 0)
	errUserNotFound      = errors.NotFound("user", 0)
	errInvalidPassword   = errors.InvalidPassword(0)
)
