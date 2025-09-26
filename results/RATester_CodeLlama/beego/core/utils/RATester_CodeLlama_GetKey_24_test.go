package utils

import (
	"fmt"
	"testing"
)

func TestGetKey_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SimpleKV{
		Key:   "key",
		Value: "value",
	}
	if s.GetKey() != "key" {
		t.Errorf("GetKey() = %v, want %v", s.GetKey(), "key")
	}
}
