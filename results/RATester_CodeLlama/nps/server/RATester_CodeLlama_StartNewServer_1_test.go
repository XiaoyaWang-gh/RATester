package server

import (
	"fmt"
	"testing"
)

func TestStartNewServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup your test
	// TODO: Exercise SUT
	// TODO: Verify results
}
