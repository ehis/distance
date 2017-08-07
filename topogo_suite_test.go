package topogo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTopogo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Topogo Suite")
}
