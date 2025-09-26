package hints

import (
	"fmt"
	"testing"
)

func TestLimit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := int64(10)
	h := Limit(d)
	if h.GetKey() != KeyLimit {
		t.Errorf("GetKey() = %v, want %v", h.GetKey(), KeyLimit)
	}
	if h.GetValue() != d {
		t.Errorf("GetValue() = %v, want %v", h.GetValue(), d)
	}
}
