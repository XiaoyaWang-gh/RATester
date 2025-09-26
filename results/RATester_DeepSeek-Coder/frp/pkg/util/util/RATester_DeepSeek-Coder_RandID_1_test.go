package util

import (
	"fmt"
	"testing"
)

func TestRandID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id, err := RandID()
	if err != nil {
		t.Errorf("RandID() error = %v", err)
		return
	}
	if len(id) != 16 {
		t.Errorf("RandID() length = %v, want 16", len(id))
	}
}
