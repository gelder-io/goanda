package goanda_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoanda(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goanda Suite")
}
