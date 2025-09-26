package client

import (
	"fmt"
	"testing"
)

func TestAdd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := make(PathParam)
	pp.Add("key", "val")
	if pp["key"] != "val" {
		t.Errorf("Expected 'val', got '%s'", pp["key"])
	}
}
