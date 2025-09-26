package e2e

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/test/e2e/framework"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestRunE2ETests_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gomega.RegisterFailHandler(framework.Fail)

	suiteConfig, reporterConfig := ginkgo.GinkgoConfiguration()

	suiteConfig.EmitSpecProgress = true

	suiteConfig.RandomizeAllSpecs = true

	log.Infof("Starting e2e run %q on Ginkgo node %d of total %d",
		framework.RunID, suiteConfig.ParallelProcess, suiteConfig.ParallelTotal)
	ginkgo.RunSpecs(t, "frp e2e suite", suiteConfig, reporterConfig)
}
