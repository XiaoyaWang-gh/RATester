package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaDash{Key: "test"}
	if a.GetKey() != "test" {
		t.Errorf("GetKey() = %v, want %v", a.GetKey(), "test")
	}
}
