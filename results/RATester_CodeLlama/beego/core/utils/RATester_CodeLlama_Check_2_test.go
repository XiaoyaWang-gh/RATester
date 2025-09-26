package utils

import (
	"fmt"
	"testing"
)

func TestCheck_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := NewBeeMap()
	m.Set("key", "value")
	if !m.Check("key") {
		t.Error("Check failed")
	}
}
