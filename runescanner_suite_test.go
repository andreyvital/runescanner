package runescanner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRunescanner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runescanner Suite")
}
