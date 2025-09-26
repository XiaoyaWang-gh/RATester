package hints

import (
	"fmt"
	"testing"
)

func TestGetValue_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Hint{
		key:   "key",
		value: "value",
	}
	if s.GetValue() != "value" {
		t.Errorf("GetValue() = %v, want %v", s.GetValue(), "value")
	}
}
