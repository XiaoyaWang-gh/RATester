package hints

import (
	"fmt"
	"testing"
)

func TestOffset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := int64(10)
	h := Offset(d)
	if h.GetKey() != KeyOffset {
		t.Errorf("Expected %v, got %v", KeyOffset, h.GetKey())
	}
	if h.GetValue() != d {
		t.Errorf("Expected %v, got %v", d, h.GetValue())
	}
}
