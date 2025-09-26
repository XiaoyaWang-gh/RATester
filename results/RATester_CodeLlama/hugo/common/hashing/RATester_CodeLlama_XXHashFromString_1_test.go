package hashing

import (
	"fmt"
	"testing"
)

func TestXXHashFromString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := "hello world"
	h, err := XXHashFromString(s)
	if err != nil {
		t.Errorf("XXHashFromString failed: %v", err)
	}
	if h != 1498963919907057248 {
		t.Errorf("XXHashFromString failed: %v", h)
	}
}
