package context

import (
	"fmt"
	"testing"
)

func TestNewOutput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := NewOutput()
	if output == nil {
		t.Error("Expected non-nil output, got nil")
	}
}
