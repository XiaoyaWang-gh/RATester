package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Phone{
		Key: "1234567890",
	}
	if p.GetKey() != "1234567890" {
		t.Errorf("GetKey() = %v, want %v", p.GetKey(), "1234567890")
	}
}
