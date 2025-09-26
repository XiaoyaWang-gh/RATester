package echo

import (
	"fmt"
	"testing"
)

func TestBools_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	binder := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "true"
		},
	}

	var dest []bool
	err := binder.Bools("param", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(dest) != 1 {
		t.Errorf("Expected dest length to be 1, but got %d", len(dest))
	}

	if !dest[0] {
		t.Errorf("Expected dest[0] to be true, but got false")
	}
}
