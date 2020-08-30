package ferrules_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFerrules(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ferrules Suite")
}
