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
		t.Errorf("RandID() returned an error: %v", err)
	}
	if len(id) != 16 {
		t.Errorf("RandID() returned an id of length %d, expected 16", len(id))
	}
}
