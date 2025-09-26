package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Enum{
		Rules: "enum",
		Key:   "key",
	}
	if e.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", e.GetKey(), "key")
	}
}
