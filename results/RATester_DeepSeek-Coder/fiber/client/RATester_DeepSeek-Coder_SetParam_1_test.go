package client

import (
	"fmt"
	"testing"
)

func TestSetParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := make(PathParam)
	pp.SetParam("key", "val")
	if pp["key"] != "val" {
		t.Errorf("Expected 'val', got '%v'", pp["key"])
	}
}
