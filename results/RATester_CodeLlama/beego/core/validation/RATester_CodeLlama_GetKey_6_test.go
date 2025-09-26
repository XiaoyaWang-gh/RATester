package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := NoMatch{
		Key: "key",
	}
	if n.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", n.GetKey(), "key")
	}
}
