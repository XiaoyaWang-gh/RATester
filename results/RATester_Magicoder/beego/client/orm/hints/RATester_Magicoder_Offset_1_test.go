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
	hint := Offset(d)

	if hint.GetKey() != KeyOffset {
		t.Errorf("Expected key to be %v, but got %v", KeyOffset, hint.GetKey())
	}

	if hint.GetValue() != d {
		t.Errorf("Expected value to be %v, but got %v", d, hint.GetValue())
	}
}
