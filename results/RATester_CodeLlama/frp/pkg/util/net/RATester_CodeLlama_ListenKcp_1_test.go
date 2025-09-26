package net

import (
	"fmt"
	"testing"
)

func TestListenKcp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
