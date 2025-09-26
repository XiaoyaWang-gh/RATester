package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := IP{Key: "key"}
	if i.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", i.GetKey(), "key")
	}
}
