package es

import (
	"fmt"
	"testing"
)

func TestWriteMsg_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup your test
	// TODO: mock your dependencies
	// TODO: run your test
	// TODO: assert what should be
}
