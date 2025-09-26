package htesting

import (
	"fmt"
	"testing"
)

func TestIsWindows_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if IsWindows() {
		t.Error("IsWindows() should return false")
	}
}
