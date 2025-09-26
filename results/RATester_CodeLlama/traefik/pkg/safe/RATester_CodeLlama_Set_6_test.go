package safe

import (
	"fmt"
	"testing"
)

func TestSet_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Safe{}
	value := "test"
	s.Set(value)
	if s.Get() != value {
		t.Errorf("expected %v, got %v", value, s.Get())
	}
}
