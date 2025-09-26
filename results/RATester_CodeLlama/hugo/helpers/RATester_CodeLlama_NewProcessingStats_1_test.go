package helpers

import (
	"fmt"
	"testing"
)

func TestNewProcessingStats_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	ps := NewProcessingStats(name)
	if ps.Name != name {
		t.Errorf("NewProcessingStats() = %v, want %v", ps.Name, name)
	}
}
