package net

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

	// TODO
}
