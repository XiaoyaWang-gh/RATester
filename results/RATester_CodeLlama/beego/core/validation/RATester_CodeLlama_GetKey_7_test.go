package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Base64{}
	b.Key = "test"
	if b.GetKey() != "test" {
		t.Errorf("GetKey() = %v, want %v", b.GetKey(), "test")
	}
}
