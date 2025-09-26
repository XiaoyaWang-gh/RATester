package util

import (
	"fmt"
	"testing"
)

func TestRandIDWithLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id, err := RandIDWithLen(10)
	if err != nil {
		t.Errorf("RandIDWithLen error: %v", err)
	}
	if len(id) != 10 {
		t.Errorf("RandIDWithLen id length error: %v", id)
	}
}
