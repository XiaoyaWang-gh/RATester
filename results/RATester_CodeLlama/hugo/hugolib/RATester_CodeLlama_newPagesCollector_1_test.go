package hugolib

import (
	"fmt"
	"testing"
)

func TestNewPagesCollector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup test
	// TODO: Exercise SUT
	// TODO: Verify results
}
