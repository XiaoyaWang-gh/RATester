package maps

import (
	"fmt"
	"testing"
)

func TestNewScratcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratcher := NewScratcher()
	if scratcher == nil {
		t.Errorf("NewScratcher() returned nil")
	}
}
