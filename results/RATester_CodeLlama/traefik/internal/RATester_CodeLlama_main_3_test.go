package main

import (
	"fmt"
	"testing"
)

func TestMain_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup your test
	// TODO: Exercise SUT
	// TODO: Verify results
}
