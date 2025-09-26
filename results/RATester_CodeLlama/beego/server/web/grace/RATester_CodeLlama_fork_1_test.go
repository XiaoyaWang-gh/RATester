package grace

import (
	"fmt"
	"testing"
)

func TestFork_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup
	// TODO: test
	// TODO: teardown
}
