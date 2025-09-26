package main

import (
	"fmt"
	"testing"
)

func TestMain_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	// when
	// then
}
