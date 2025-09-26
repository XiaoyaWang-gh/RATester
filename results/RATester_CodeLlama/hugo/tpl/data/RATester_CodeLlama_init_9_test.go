package data

import (
	"fmt"
	"testing"
)

func TestInit_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup the test case
	// TODO: Test the result
}
