package utils

import (
	"fmt"
	"testing"
)

func TestSelfPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := SelfPath()
	if path == "" {
		t.Errorf("SelfPath() failed")
	}
}
