package client

import (
	"fmt"
	"testing"
)

func Testinit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := New()
	if client == nil {
		t.Error("Expected client to be initialized, but it was nil")
	}
}
