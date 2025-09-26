package utils

import (
	"fmt"
	"testing"
)

func TestGetValue_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &SimpleKV{
		Key:   "key",
		Value: "value",
	}
	if s.GetValue() != "value" {
		t.Errorf("GetValue() = %v, want %v", s.GetValue(), "value")
	}
}
