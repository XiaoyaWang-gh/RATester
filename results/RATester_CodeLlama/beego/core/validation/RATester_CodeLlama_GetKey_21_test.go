package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := Length{Key: "key"}
	if l.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", l.GetKey(), "key")
	}
}
