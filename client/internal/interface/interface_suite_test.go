package _interface_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInterface(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interface Suite")
}
