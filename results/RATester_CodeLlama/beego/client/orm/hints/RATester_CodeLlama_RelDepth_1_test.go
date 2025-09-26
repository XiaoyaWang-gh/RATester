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

	d := 1
	h := RelDepth(d)
	if h.GetKey() != KeyRelDepth {
		t.Errorf("Expected %v, got %v", KeyRelDepth, h.GetKey())
	}
	if h.GetValue() != d {
		t.Errorf("Expected %v, got %v", d, h.GetValue())
	}
}
