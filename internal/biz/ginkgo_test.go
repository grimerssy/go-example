package biz

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_UseCases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Business logic suite")
}
