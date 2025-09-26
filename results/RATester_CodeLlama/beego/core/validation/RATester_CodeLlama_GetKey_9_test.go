package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mobile{}
	m.Key = "key"
	if m.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", m.GetKey(), "key")
	}
}
