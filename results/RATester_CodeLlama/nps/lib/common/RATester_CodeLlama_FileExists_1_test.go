package common

import (
	"fmt"
	"testing"
)

func TestFileExists_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if FileExists("test.txt") {
		t.Errorf("FileExists() = true, want false")
	}
}
