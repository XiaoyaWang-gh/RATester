package hugolib

import (
	"fmt"
	"testing"
)

func TestisMultilingual_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{
		Sites: []*Site{
			{},
			{},
		},
	}

	result := h.isMultilingual()

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
