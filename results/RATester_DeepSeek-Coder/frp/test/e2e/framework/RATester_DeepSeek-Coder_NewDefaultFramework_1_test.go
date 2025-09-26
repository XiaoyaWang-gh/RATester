package framework

import (
	"fmt"
	"testing"

	"github.com/onsi/ginkgo"
)

func TestNewDefaultFramework_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ginkgo.RunSpecs(t, "NewDefaultFramework Suite")
}
