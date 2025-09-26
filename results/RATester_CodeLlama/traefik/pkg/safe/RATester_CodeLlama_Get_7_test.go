package safe

import (
	"fmt"
	"testing"
)

func TestGet_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Safe{value: "test"}
	if s.Get() != "test" {
		t.Errorf("Get() = %v, want %v", s.Get(), "test")
	}
}
