package images

import (
	"fmt"
	"testing"
)

func TestString_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := resamp{name: "test"}
	if r.String() != "test" {
		t.Errorf("String() = %v, want %v", r.String(), "test")
	}
}
