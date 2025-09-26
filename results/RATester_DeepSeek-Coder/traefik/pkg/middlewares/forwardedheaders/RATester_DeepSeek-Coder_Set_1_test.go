package forwardedheaders

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := make(unsafeHeader)
	h.Set("Content-Type", "application/json")

	v, ok := h["Content-Type"]
	if !ok {
		t.Errorf("Expected key 'Content-Type' to exist")
	}

	if v[0] != "application/json" {
		t.Errorf("Expected value 'application/json', got '%s'", v[0])
	}
}
