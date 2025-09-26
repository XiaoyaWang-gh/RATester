package e2e

import (
	"fmt"
	"testing"
)

func TestCleanupSuite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	CleanupSuite()
	if t.Failed() {
		t.Errorf("Test failed")
	}
}
