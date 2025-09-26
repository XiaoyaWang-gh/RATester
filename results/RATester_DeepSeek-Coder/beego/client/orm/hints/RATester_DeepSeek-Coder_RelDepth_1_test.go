package hints

import (
	"fmt"
	"testing"
)

func TestRelDepth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := 5
	hint := RelDepth(d)

	if hint.GetKey() != KeyRelDepth {
		t.Errorf("Expected key %v, got %v", KeyRelDepth, hint.GetKey())
	}

	if hint.GetValue() != d {
		t.Errorf("Expected value %v, got %v", d, hint.GetValue())
	}
}
