package main

import (
	"fmt"
	"testing"
)

func TestMain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup your test
	// TODO: Test
	// TODO: Teardown your test
}
