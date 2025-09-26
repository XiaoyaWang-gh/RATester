package utils

import (
	"fmt"
	"testing"
)

func TestGrepFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	patten := "test"
	filename := "test.txt"
	lines, err := GrepFile(patten, filename)
	if err != nil {
		t.Errorf("GrepFile error: %v", err)
	}
	if len(lines) != 1 {
		t.Errorf("GrepFile error: %v", lines)
	}
}
