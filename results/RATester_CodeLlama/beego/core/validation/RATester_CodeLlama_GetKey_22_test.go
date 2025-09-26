package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaNumeric{Key: "key"}
	if a.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", a.GetKey(), "key")
	}
}
