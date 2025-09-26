package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	z := ZipCode{
		Key: "12345",
	}
	if z.GetKey() != "12345" {
		t.Errorf("GetKey() = %v, want %v", z.GetKey(), "12345")
	}
}
