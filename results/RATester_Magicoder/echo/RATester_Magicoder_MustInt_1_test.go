package echo

import (
	"fmt"
	"testing"
)

func TestMustInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "10"
		},
	}

	var dest int
	err := b.MustInt("test", &dest)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if dest != 10 {
		t.Errorf("Expected dest to be 10, but got %d", dest)
	}
}
