package plugins

import (
	"fmt"
	"testing"
)

func TestComputeHash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filepath := "test.txt"
	expectedHash := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd154708f8b3d39d23"

	hash, err := computeHash(filepath)
	if err != nil {
		t.Errorf("computeHash() error = %v", err)
		return
	}

	if hash != expectedHash {
		t.Errorf("computeHash() = %v, want %v", hash, expectedHash)
	}
}
