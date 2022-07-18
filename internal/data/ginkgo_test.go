package data

import (
	"testing"

	"github.com/grimerssy/go-example/pkg/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data access suite")
}

var (
	errUserAlreadyExists = errors.AlreadyExists("user", 0)
	errUserNotFound      = errors.NotFound("user", 0)
)
