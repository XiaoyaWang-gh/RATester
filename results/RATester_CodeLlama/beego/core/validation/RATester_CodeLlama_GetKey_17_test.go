package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := MaxSize{
		Max: 10,
		Key: "key",
	}
	if m.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", m.GetKey(), "key")
	}
}
