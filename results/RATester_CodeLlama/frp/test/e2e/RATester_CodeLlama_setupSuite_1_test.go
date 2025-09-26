package e2e

import (
	"fmt"
	"testing"
)

func TestSetupSuite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	setupSuite()
}
