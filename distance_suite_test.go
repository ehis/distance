package distance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDistance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Distance Suite")
}
