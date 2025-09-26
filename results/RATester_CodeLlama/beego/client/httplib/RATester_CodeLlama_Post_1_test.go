package httplib

import (
	"fmt"
	"testing"
)

func TestPost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup the test case
	// TODO: Call the method to test
	// TODO: Verify the results
}
