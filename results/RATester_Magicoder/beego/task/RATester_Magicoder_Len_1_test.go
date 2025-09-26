package task

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ms := &MapSorter{
		Keys: []string{"a", "b", "c"},
	}

	if ms.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", ms.Len())
	}
}
