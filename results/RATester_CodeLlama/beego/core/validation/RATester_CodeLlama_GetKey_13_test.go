package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Required{}
	r.Key = "key"
	if r.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", r.GetKey(), "key")
	}
}
