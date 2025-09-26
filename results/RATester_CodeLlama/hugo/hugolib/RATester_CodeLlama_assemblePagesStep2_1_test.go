package hugolib

import (
	"fmt"
	"testing"
)

func TestAssemblePagesStep2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: This test case is incomplete.
}
