package api

import (
	"fmt"
	"testing"
)

func TestInit_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Your test code here
	t.Log("Testing init")
}
