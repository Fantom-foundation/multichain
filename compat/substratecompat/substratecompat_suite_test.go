package substratecompat_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSubstratecompat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Substratecompat Suite")
}
