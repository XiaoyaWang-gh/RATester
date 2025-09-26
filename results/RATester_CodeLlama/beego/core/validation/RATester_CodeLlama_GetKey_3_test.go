package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := Numeric{Key: "key"}
	if n.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", n.GetKey(), "key")
	}
}
